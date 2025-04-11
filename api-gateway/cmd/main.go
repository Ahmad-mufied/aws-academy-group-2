package main

import (
	"flag"
	"net"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"

	"github.com/Ahmad-mufied/aws-academy-group-2/api-gateway/internal/breaker"
	"github.com/Ahmad-mufied/aws-academy-group-2/api-gateway/internal/config"
	"github.com/Ahmad-mufied/aws-academy-group-2/api-gateway/internal/lb"
	"github.com/Ahmad-mufied/aws-academy-group-2/api-gateway/internal/proxy"
)

func main() {
	cfgPath := flag.String("config", "api-config.toml", "path to TOML config")
	flag.Parse()

	cfg := config.MustLoad(*cfgPath)

	// Build RoundRobin and CircuitBreaker maps
	rrMap := make(map[string]*lb.RoundRobin)
	cbMap := make(map[string]*breaker.Breaker)
	for _, svc := range cfg.Services {
		rrMap[svc.Name] = lb.New([]string{svc.URL})
		cbMap[svc.Name] = breaker.New(svc.Name)
	}

	// HTTP client: disable keep‑alives for per‑request DNS
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
		},
		Timeout: 30 * time.Second,
	}

	e := echo.New()
	e.HideBanner = true

	// Global rate limiter: 100 req/min
	limiter := rate.NewLimiter(rate.Every(time.Minute/100), 100)
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !limiter.Allow() {
				return c.String(http.StatusTooManyRequests, "Rate limit exceeded")
			}
			return next(c)
		}
	})

	// Dynamically mount each service route
	for _, svc := range cfg.Services {
		rr := rrMap[svc.Name]
		cb := cbMap[svc.Name]

		group := e.Group("/" + svc.Name)
		group.Any("/*", func(c echo.Context) error {
			target := rr.Next()
			// Circuit-breaker wrap
			resp, err := cb.Execute(func() (*http.Response, error) {
				u := target + c.Request().RequestURI
				req, _ := http.NewRequest(c.Request().Method, u, c.Request().Body)
				req.Header = c.Request().Header
				return client.Do(req)
			})
			if err != nil {
				return c.String(http.StatusServiceUnavailable, "Service unavailable")
			}
			defer resp.Body.Close()
			return proxy.ProxyRequest(target, c, client)
		})
	}

	e.Logger.Fatal(e.Start(":8090"))
}

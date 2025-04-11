package proxy

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
)

// ProxyRequest forwards the incoming Echo context to the target URL.
func ProxyRequest(target string, c echo.Context, client *http.Client) error {
	newPath := strings.TrimPrefix(c.Request().URL.Path, "/"+strings.Split(c.Path(), "/")[1])
	c.Request().URL.Path = newPath
	u, err := url.Parse(target)
	if err != nil {
		return err
	}
	backendURL := u.ResolveReference(c.Request().URL)

	req, err := http.NewRequest(c.Request().Method, backendURL.String(), c.Request().Body)
	if err != nil {
		return err
	}
	req.Header = c.Request().Header

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	for k, vv := range resp.Header {
		for _, v := range vv {
			c.Response().Header().Add(k, v)
		}
	}
	c.Response().WriteHeader(resp.StatusCode)
	_, err = io.Copy(c.Response().Writer, resp.Body)
	return err
}

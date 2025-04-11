package breaker

import (
	"errors"
	"net/http"
	"time"

	"github.com/sony/gobreaker"
)

// Breaker wraps a gobreaker.CircuitBreaker.
type Breaker struct {
	cb *gobreaker.CircuitBreaker
}

// New constructs a breaker with sensible defaults.
func New(name string) *Breaker {
	settings := gobreaker.Settings{
		Name:        name,
		MaxRequests: 3,
		Interval:    60 * time.Second,
		Timeout:     30 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures > 5
		},
	}
	return &Breaker{cb: gobreaker.NewCircuitBreaker(settings)}
}

// Execute runs the provided request function through the circuit breaker.
func (b *Breaker) Execute(reqFunc func() (*http.Response, error)) (*http.Response, error) {
	result, err := b.cb.Execute(func() (interface{}, error) {
		return reqFunc()
	})
	if err != nil {
		return nil, err
	}
	resp, ok := result.(*http.Response)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return resp, nil
}

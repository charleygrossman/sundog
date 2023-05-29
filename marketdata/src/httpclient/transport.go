package httpclient

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// NOTE: RoundTrippers are invoked in the order they're provided.
func NewTransportWithOptions(opts ...TransportOptionFunc) http.RoundTripper {
	t := http.DefaultTransport
	for i := len(opts) - 1; i >= 0; i-- {
		t = opts[i](t)
	}
	return t
}

type TransportOptionFunc func(http.RoundTripper) http.RoundTripper

type BearerAuthRoundTripper struct {
	next   http.RoundTripper
	apiKey string
}

func WithBearerAuth(apiKey string) TransportOptionFunc {
	return func(next http.RoundTripper) http.RoundTripper {
		return &BearerAuthRoundTripper{
			next:   next,
			apiKey: apiKey,
		}
	}
}

func (t *BearerAuthRoundTripper) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.apiKey))
	return t.next.RoundTrip(req)
}

type DelayRoundTripper struct {
	next  http.RoundTripper
	delay time.Duration
}

func WithDelay(delay time.Duration) TransportOptionFunc {
	return func(next http.RoundTripper) http.RoundTripper {
		return &DelayRoundTripper{
			next:  next,
			delay: delay,
		}
	}
}

func (t *DelayRoundTripper) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	time.Sleep(t.delay)
	return t.next.RoundTrip(req)
}

type RequestLoggerRoundTripper struct {
	next   http.RoundTripper
	logger *log.Logger
}

func WithRequestLogger(logger *log.Logger) TransportOptionFunc {
	return func(next http.RoundTripper) http.RoundTripper {
		return &RequestLoggerRoundTripper{
			next:   next,
			logger: logger,
		}
	}
}

func (t *RequestLoggerRoundTripper) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	t.logger.Println(fmt.Sprintf("request: method=%s url=%s", req.Method, req.URL))
	return t.next.RoundTrip(req)
}

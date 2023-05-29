package httpclient

import (
	"net/http"
	"time"
)

func NewClient(opts ...ClientOptionFunc) *http.Client {
	c := new(http.Client)
	for _, opt := range opts {
		opt(c)
	}
	return c
}

type ClientOptionFunc func(client *http.Client)

func WithTimeout(timeout *time.Duration) ClientOptionFunc {
	return func(c *http.Client) {
		c.Timeout = *timeout
	}
}

func WithTransport(transport http.RoundTripper) ClientOptionFunc {
	return func(c *http.Client) {
		c.Transport = transport
	}
}

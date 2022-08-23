package tinybird

import (
	"net/http"
	"time"
)

// Custom HTTP client for this module.
var Client HTTPClient

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Initialize module.
func init() {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxIdleConns = 100
	transport.MaxConnsPerHost = 100
	transport.MaxIdleConnsPerHost = 100

	Client = &http.Client{
		Timeout:   time.Duration(30) * time.Second,
		Transport: transport,
	}
}

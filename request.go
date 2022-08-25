package tinybird

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Request struct {
	Elapsed  Duration // Elapsed time of client request.
	Error    error    // Error on client request.
	Method   string   // Define HTTP method.
	Pipe     Pipe     // Pipe details.
	Response Response // Response data.
}

// Custom HTTP client for this module.
var Client HTTPClient

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Initialize module.
func init() {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxIdleConns = MAX_IDLE_CONNS
	transport.MaxConnsPerHost = MAX_CONNS_PER_HOST
	transport.MaxIdleConnsPerHost = MAX_IDLE_CONNS_PER_HOST

	Client = &http.Client{
		Timeout:   time.Duration(CONNS_TIMEOUT) * time.Second,
		Transport: transport,
	}
}

// Execute request.
func (r *Request) Execute() error {
	r.Error = r.Elapsed.Do(func() (err error) {
		req, err := r.newRequest()
		if err != nil {
			return err
		}

		res, err := Client.Do(req)
		if err != nil {
			return err
		}

		return r.readBody(res)
	})

	return r.Error
}

// Create new request.
func (r *Request) newRequest() (*http.Request, error) {
	req, err := http.NewRequest(r.Method, r.Pipe.GetURL(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", r.Pipe.Workspace.Token))
	req.URL.RawQuery = r.Pipe.Parameters.Encode()

	return req, nil
}

// Read response body from request.
func (r *Request) readBody(resp *http.Response) (err error) {
	defer resp.Body.Close()

	r.Response.Status = resp.StatusCode
	r.Response.Body, err = io.ReadAll(resp.Body)
	r.Response.Decode()

	return err
}

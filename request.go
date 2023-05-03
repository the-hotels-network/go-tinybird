package tinybird

import (
	"fmt"
	"net/http"
	"time"

	"github.com/the-hotels-network/go-tinybird/internal/env"

	log "github.com/sirupsen/logrus"
)

// Basic request struct.
type Request struct {
	Elapsed              Duration // Elapsed time of client request.
	Error                error    // Error on client request.
	Method               string   // Define HTTP method.
	Pipe                 Pipe     // Pipe details.
	Response             Response // Response data.
	NewLineDelimitedJSON bool     // Enable NewLine-Delimited JSON.
}

// Custom HTTP client for this module.
var Client HTTPClient

// HTTPClient interface.
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
	req, err := http.NewRequest(r.Method, r.URL(), nil)
	if err != nil {
		return nil, err
	}

	log.WithFields(log.Fields{
		"http-message-type": "request",
		"uri":               r.URI(),
	}).Debug("tinybird")

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", r.Pipe.Workspace.Token))
	req.URL.RawQuery = r.Pipe.Parameters.Encode()

	return req, nil
}

// Read response body from request.
func (r *Request) readBody(resp *http.Response) (err error) {
	defer resp.Body.Close()

	r.Response.NewLineDelimitedJSON = r.NewLineDelimitedJSON
	r.Response.Status = resp.StatusCode
	r.Response.Body = resp.Body
	err = r.Response.Decode()

	log.WithFields(log.Fields{
		"http-message-type": "response",
		"uri":               r.URI(),
		"status":            r.Response.Status,
	}).Debug("tinybird")

	return err
}

// Build and return the pipe URL.
func (r *Request) URL() string {
	var baseUrl string
	if r.Pipe.URL != "" {
		baseUrl = r.Pipe.URL
	} else {
		baseUrl = URL_BASE
	}

	return fmt.Sprintf(
		"%s/%s.%s",
		baseUrl,
		r.Pipe.Name,
		r.Format(),
	)
}

// Verify the variable NewLineDelimitedJSON  value to return json or ndjson.
func (r *Request) Format() string {
	if r.NewLineDelimitedJSON {
		return "ndjson"
	}

	if env.GetBool("TB_NDJSON", false) {
		return "ndjson"
	}

	return "json"
}

// Return concatened URL and Query String to generate a URI.
func (r *Request) URI() string {
	return fmt.Sprintf("%s?%s", r.URL(), r.Pipe.Parameters.Encode())
}

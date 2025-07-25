package tinybird

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/the-hotels-network/go-tinybird/internal/env"

	log "github.com/sirupsen/logrus"
)

const (
	Format string = "json"
	JSON          = "json"
	NDJSON        = "ndjson"
	CSV           = "csv"
)

// Basic request struct.
type Request struct {
	Elapsed  Duration            // Elapsed time of client request.
	Error    error               // Error on client request.
	Method   string              // Define HTTP method.
	Pipe     Pipe                // Pipe details.
	Response Response            // Response data.
	Format   string              // Return format.
	Before   func(*Request) bool // Run before execute request.
	After    func(*Request)      // Run after execute request.
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
	if r.Before == nil {
		r.Before = func(*Request) bool { return false }
	}
	if r.After == nil {
		r.After = func(*Request) {}
	}

	r.Error = r.Elapsed.Do(func() (err error) {
		if !r.Before(r) {
			req, err := r.newRequest()
			if err != nil {
				return err
			}

			res, err := Client.Do(req)
			if err != nil {
				return err
			}

			err = r.readBody(res)
		}
		r.After(r)
		return err
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
	req.URL.RawQuery = r.Pipe.GetParameters()

	return req, nil
}

// Read response body from request.
func (r *Request) readBody(resp *http.Response) (err error) {
	defer resp.Body.Close()

	r.Response.Format = r.Format
	r.Response.Status = resp.StatusCode
	r.Response.Header = resp.Header
	r.Response.Raw = resp.Body
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
		r.GetFormat(),
	)
}

// Verify the variable Format value to return json, ndjson or csv.
func (r *Request) GetFormat() string {
	if r.Format == JSON || r.Format == NDJSON || r.Format == CSV {
		return r.Format
	}

	if env.GetBool("TB_NDJSON", false) {
		return NDJSON
	}

	return JSON
}

// Return concatened URL and Query String to generate a URI.
func (r *Request) URI() string {
	qs, _ := url.QueryUnescape(r.Pipe.GetParameters())
	return fmt.Sprintf("%s?%s", r.URL(), qs)
}

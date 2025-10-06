package tinybird

import (
	"bytes"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// Basic request struct.
type Request struct {
	After    func(*Request)      // Run after execute request.
	Before   func(*Request) bool // Run before execute request.
	Data     []byte              // Data to send.
	Elapsed  Duration            // Elapsed time of client request.
	Error    error               // Error on client request.
	Event    Event               // Send event to data source.
	Method   string              // Define HTTP method.
	Pipe     Pipe                // Pipe details.
	Response Response            // Response data.
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
	req, err := http.NewRequest(r.Method, r.URI(), bytes.NewBuffer(r.Data))
	if err != nil {
		return nil, err
	}

	log.WithFields(log.Fields{
		"http-message-type": "request",
		"uri":               r.URI(),
	}).Debug("tinybird")

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", r.Token())

	if r.Pipe.Workspace.IsSet() {
		req.URL.RawQuery = r.Pipe.GetParameters()
	}

	return req, nil
}

// Read response body from request.
func (r *Request) readBody(resp *http.Response) (err error) {
	defer resp.Body.Close()

	r.Response.Format = r.Format()
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

func (r *Request) URL() (out string) {
	if r.Pipe.Workspace.IsSet() {
		out = r.Pipe.GetURL()
	} else if r.Event.Workspace.IsSet() {
		out = r.Event.GetURL()
	}

	return out
}

func (r *Request) Token() (out string) {
	if r.Pipe.Workspace.IsSet() {
		out = r.Pipe.Workspace.GetToken()
	} else if r.Event.Workspace.IsSet() {
		out = r.Event.Workspace.GetToken()
	}

	return out
}

func (r *Request) URI() (out string) {
	if r.Pipe.Workspace.IsSet() {
		out = r.Pipe.GetURI()
	} else if r.Event.Workspace.IsSet() {
		out = r.Event.GetURI()
	}

	return out
}

func (r *Request) Format() (out string) {
	if r.Pipe.Workspace.IsSet() {
		out = r.Pipe.GetFormat()
	}

	return out
}

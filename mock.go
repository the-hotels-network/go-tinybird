package tinybird

import (
	"bytes"
	"io"
	"net/http"
)

type MockRoundTripper func(r *http.Request) *http.Response

func InjectHTTPClient(httpClient *http.Client) {
	Client = httpClient
}

func (f MockRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r), nil
}

func MockResponse(statusCode int, body string) {
	InjectHTTPClient(&http.Client{
		Transport: MockRoundTripper(func(r *http.Request) *http.Response {
			return &http.Response{
				StatusCode: statusCode,
				Body:       io.NopCloser(bytes.NewReader([]byte(body))),
			}
		}),
	})
}

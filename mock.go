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
	MockResponseWithRequestInspect(statusCode, body, nil)
}

func MockResponseWithRequestInspect(statusCode int, body string, requestInspect func(r *http.Request)) {
	InjectHTTPClient(&http.Client{
		Transport: MockRoundTripper(func(r *http.Request) *http.Response {
			if requestInspect != nil {
				requestInspect(r)
			}

			return &http.Response{
				StatusCode: statusCode,
				Body:       io.NopCloser(bytes.NewReader([]byte(body))),
			}
		}),
	})
}

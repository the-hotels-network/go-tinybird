package tinybird

import (
	"bytes"
	"io"
	"net/http"
)

type MockRoundTripper func(r *http.Request) (*http.Response, error)

func (f MockRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func MockResponse(statusCode int, body string, requestInspect func(r *http.Request)) {
	Client = &http.Client{
		Transport: MockRoundTripper(func(r *http.Request) (*http.Response, error) {
			if requestInspect != nil {
				requestInspect(r)
			}

			return &http.Response{
				StatusCode:    statusCode,
				Body:          io.NopCloser(bytes.NewReader([]byte(body))),
				ContentLength: int64(len([]byte(body))),
				Request:       r,
			}, nil
		}),
	}
}

func MockResponseFunc(statusCode int,	bodyFn func(string) string,	requestInspect func(r *http.Request)) {
	Client = &http.Client{
		Transport: MockRoundTripper(func(r *http.Request) (*http.Response, error) {
			if requestInspect != nil {
				requestInspect(r)
			}

			body := bodyFn(r.URL.String())
			return &http.Response{
				StatusCode:    statusCode,
				Body:          io.NopCloser(bytes.NewReader([]byte(body))),
				ContentLength: int64(len([]byte(body))),
				Request:       r,
			}, nil
		}),
	}
}

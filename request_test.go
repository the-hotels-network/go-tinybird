package tinybird_test

import (
	"net/http"
	"testing"

	"github.com/the-hotels-network/go-tinybird"

	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	req := tinybird.Request{
		Method: http.MethodGet,
		Pipe: tinybird.Pipe{
			Name: "test",
			Workspace: tinybird.Workspace{
				Name: "test",
			},
		},
	}

	assert.Equal(t, req.URL(), "https://api.tinybird.co/v0/pipes/test.json")
}

func TestRequestNDJSON(t *testing.T) {
	req := tinybird.Request{
		Method: http.MethodGet,
		Pipe: tinybird.Pipe{
			Name: "test",
			Workspace: tinybird.Workspace{
				Name: "test",
			},
		},
		NewLineDelimitedJSON: true,
	}

	assert.Equal(t, req.URL(), "https://api.tinybird.co/v0/pipes/test.ndjson")
}

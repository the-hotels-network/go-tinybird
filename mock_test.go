package tinybird_test

import (
	"net/http"
	"testing"

	"github.com/the-hotels-network/go-tinybird"

	"github.com/stretchr/testify/assert"
)

func TestMock(t *testing.T) {
	req := tinybird.Request{
		Method: http.MethodGet,
		Pipe: tinybird.Pipe{
			Name:       "test",
			Workspace: tinybird.Workspace{
				Name:  "test",
				Token: "testoken",
			},
		},
	}

	tinybird.MockResponse(
		http.StatusOK,
		`ok`,
		func(r *http.Request) {
			assert.Equal(t, r.URL.String(), "https://api.tinybird.co/v0/pipes/test.json")
		},
	)

	req.Execute()
	res := req.Response

	assert.Nil(t, req.Error)
	assert.Equal(t, res.Status, http.StatusOK)
	assert.Equal(t, res.Body, "ok")
}

package tinybird_test

import (
	"net/http"
	"net/url"
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

	tinybird.MockResponse(
		http.StatusOK,
		`{"data":[{"Col1": "1", "Col2": 2}],"rows":1,"statistics":{"elapsed":0.00091042,"rows_read": 4,"bytes_read": 296}}`,
	)

	req.Execute()
	res := req.Response

	assert.Equal(t, req.URL(), "https://api.tinybird.co/v0/pipes/test.json")
	assert.Nil(t, req.Error)
	assert.Equal(t, res.Status, http.StatusOK)
	assert.Equal(t, res.Rows, uint(1))
	assert.Equal(t, res.Data, []tinybird.Row{{"Col1": "1", "Col2": float64(2)}})
}

func TestRequestWithCustomURL(t *testing.T) {
	req := tinybird.Request{
		Method: http.MethodGet,
		Pipe: tinybird.Pipe{
			Name: "test",
			URL:  "https://api.us-east.tinybird.co/v0/pipes",
			Workspace: tinybird.Workspace{
				Name: "test",
			},
		},
	}

	tinybird.MockResponse(
		http.StatusOK,
		`{"data":[{"Col1": "1", "Col2": 2}],"rows":1,"statistics":{"elapsed":0.00091042,"rows_read": 4,"bytes_read": 296}}`,
	)

	req.Execute()
	res := req.Response

	assert.Equal(t, req.URL(), "https://api.us-east.tinybird.co/v0/pipes/test.json")
	assert.Nil(t, req.Error)
	assert.Equal(t, res.Status, http.StatusOK)
	assert.Equal(t, res.Rows, uint(1))
	assert.Equal(t, res.Data, []tinybird.Row{{"Col1": "1", "Col2": float64(2)}})
}

func TestRequestWithRequestParamInspect(t *testing.T) {
	params := url.Values{}
	params.Add("start_date", "2023-05-01")
	params.Add("end_date", "2023-05-31")
	params.Add("currency", "EUR")

	req := tinybird.Request{
		Method: http.MethodGet,
		Pipe: tinybird.Pipe{
			Name:       "test",
			Parameters: params,
			Workspace: tinybird.Workspace{
				Name: "test",
			},
		},
	}

	tinybird.MockResponseWithRequestInspect(
		http.StatusOK,
		`{"data":[{"Col1": "1", "Col2": 2}],"rows":1,"statistics":{"elapsed":0.00091042,"rows_read": 4,"bytes_read": 296}}`,
		func(r *http.Request) {
			assert.Contains(t, r.URL.String(), "start_date=2023-05-01")
			assert.Contains(t, r.URL.String(), "end_date=2023-05-31")
			assert.Contains(t, r.URL.String(), "currency=EUR")
		},
	)

	req.Execute()
	res := req.Response

	assert.Nil(t, req.Error)
	assert.Equal(t, res.Status, http.StatusOK)
}

func TestNDJSON(t *testing.T) {
	r := tinybird.Request{}

	// Case 1
	r = tinybird.Request{
		Pipe: tinybird.Pipe{
			Name: "test",
		},
	}

	assert.Equal(t, r.Format(), "json")

	// Case 2
	t.Setenv("TB_NDJSON", "true")

	r = tinybird.Request{
		Pipe: tinybird.Pipe{
			Name: "test",
		},
	}

	assert.Equal(t, r.Format(), "json")

	// Case 3
	t.Setenv("TB_NDJSON", "false")

	r = tinybird.Request{
		Pipe: tinybird.Pipe{
			Name: "test",
		},
	}

	assert.Equal(t, r.Format(), "json")

	// Case 4
	t.Setenv("TB_NDJSON", "false")

	r = tinybird.Request{
		NewLineDelimitedJSON: true,
		Pipe: tinybird.Pipe{
			Name: "test",
		},
	}

	assert.Equal(t, r.Format(), "ndjson")

	// Case 5
	t.Setenv("TB_NDJSON", "true")

	r = tinybird.Request{
		NewLineDelimitedJSON: false,
		Pipe: tinybird.Pipe{
			Name: "test",
		},
	}

	assert.Equal(t, r.Format(), "json")
}

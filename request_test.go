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
			Name:  "ep_test",
			Alias: "test",
			Workspace: tinybird.Workspace{
				Name:  "test",
				Token: "testoken",
			},
		},
	}

	tinybird.MockResponse(
		http.StatusOK,
		`{"data":[{"Col1": "1", "Col2": 2}],"rows":1,"statistics":{"elapsed":0.00091042,"rows_read": 4,"bytes_read": 296}}`,
	)

	req.Execute()
	res := req.Response

	assert.Equal(t, req.URL(), "https://api.tinybird.co")
	assert.Equal(t, req.URI(), "https://api.tinybird.co/v0/pipes/ep_test.json")
	assert.Nil(t, req.Error)
	assert.Equal(t, res.Status, http.StatusOK)
	assert.Equal(t, res.Rows, uint(1))
	assert.Equal(t, res.Data, []tinybird.Row{{"Col1": "1", "Col2": float64(2)}})
}

func TestRequestWithCustomURL(t *testing.T) {
	req := tinybird.Request{
		Method: http.MethodGet,
		Pipe: tinybird.Pipe{
			Name:  "ep_test",
			Alias: "test",
			URL:   "https://api.us-east.tinybird.co",
			Workspace: tinybird.Workspace{
				Name:  "test",
				Token: "testoken",
			},
		},
	}

	tinybird.MockResponse(
		http.StatusOK,
		`{"data":[{"Col1": "1", "Col2": 2}],"rows":1,"statistics":{"elapsed":0.00091042,"rows_read": 4,"bytes_read": 296}}`,
	)

	req.Execute()
	res := req.Response

	assert.Equal(t, req.URL(), "https://api.us-east.tinybird.co")
	assert.Equal(t, req.URI(), "https://api.us-east.tinybird.co/v0/pipes/ep_test.json")
	assert.Nil(t, req.Error)
	assert.Equal(t, res.Status, http.StatusOK)
	assert.Equal(t, res.Rows, uint(1))
	assert.Equal(t, res.Data, []tinybird.Row{{"Col1": "1", "Col2": float64(2)}})
}

func TestRequestWithRequestParamInspect(t *testing.T) {
	params := url.Values{}
	params.Add("start_date", "2023-05-01")
	params.Add("request_at", "2023-05-01 23:59:59")
	params.Add("end_date", "2023-05-31")
	params.Add("currency", "EUR")

	req := tinybird.Request{
		Method: http.MethodGet,
		Pipe: tinybird.Pipe{
			Name:       "test",
			Parameters: params,
			Workspace: tinybird.Workspace{
				Name:  "test",
				Token: "testoken",
			},
		},
	}

	tinybird.MockResponseWithRequestInspect(
		http.StatusOK,
		`{"data":[{"Col1": "1", "Col2": 2}],"rows":1,"statistics":{"elapsed":0.00091042,"rows_read": 4,"bytes_read": 296}}`,
		func(r *http.Request) {
			assert.Contains(t, r.URL.String(), "start_date=2023-05-01")
			assert.Contains(t, r.URL.String(), "request_at=2023-05-01%2023%3A59%3A59")
			assert.Contains(t, r.URL.String(), "end_date=2023-05-31")
			assert.Contains(t, r.URL.String(), "currency=EUR")
		},
	)

	req.Execute()
	res := req.Response

	assert.Nil(t, req.Error)
	assert.Equal(t, res.Status, http.StatusOK)
}

func TestGetFormat(t *testing.T) {
	tests := []struct {
		name           string
		format         string
		envTB_NDJSON   string
		expectedFormat string
	}{
		{
			"format empty & TB_NDJSON=false",
			"",
			"false",
			"json",
		},
		{
			"format empty & TB_NDJSON=true",
			"",
			"true",
			"ndjson",
		},
		{
			"format=json & TB_NDJSON=false",
			"json",
			"false",
			"json",
		},
		{
			"format=json & TB_NDJSON=true",
			"json",
			"true",
			"json",
		},
		{
			"format=ndjson & TB_NDJSON=false",
			"ndjson",
			"false",
			"ndjson",
		},
		{
			"format=ndjson & TB_NDJSON=true",
			"ndjson",
			"true",
			"ndjson",
		},
		{
			"format=csv & TB_NDJSON=false",
			"csv",
			"false",
			"csv",
		},
		{
			"format=csv & TB_NDJSON=true",
			"csv",
			"true",
			"csv",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("TB_NDJSON", tc.envTB_NDJSON)

			r := tinybird.Request{
				Pipe: tinybird.Pipe{
					Format: tc.format,
					Name:   "test",
				},
			}

			assert.Equal(t, tc.expectedFormat, r.Pipe.GetFormat())
		})
	}
}

func TestRequestEvent(t *testing.T) {
	req := tinybird.Request{
		Method: http.MethodPost,
		Event: tinybird.Event{
			Datasource: "ds_test",
			Workspace: tinybird.Workspace{
				Name:  "test",
				Token: "testoken",
			},
		},
	}

	tinybird.MockResponse(
		http.StatusOK,
		`{"successful_rows":1,"quarantined_rows":0}`,
	)

	req.Execute()
	res := req.Response

	assert.Equal(t, req.URL(), "https://api.tinybird.co")
	assert.Equal(t, req.URI(), "https://api.tinybird.co/v0/events?name=ds_test")
	assert.Nil(t, req.Error)
	assert.Equal(t, res.Status, http.StatusOK)
	assert.Equal(t, res.Rows, uint(0))
	assert.Equal(t, res.Data, []tinybird.Row(nil))
}

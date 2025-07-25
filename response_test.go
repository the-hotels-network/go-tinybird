package tinybird_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/the-hotels-network/go-tinybird"

	"github.com/stretchr/testify/assert"
)

func TestJSON(t *testing.T) {
	res := tinybird.Response{}
	res.Raw = io.NopCloser(strings.NewReader(`{
    "meta": [
        {
            "name": "val1",
            "type": "String"
        },
        {
            "name": "val2",
            "type": "UInt64"
        },
        {
            "name": "val3",
            "type": "String"
        },
        {
            "name": "val4",
            "type": "UInt64"
        },
        {
            "name": "val5",
            "type": "UInt8"
        },
        {
            "name": "val6",
            "type": "Nullable"
        },
        {
            "name": "val7",
            "type": "Float64"
        }
    ],
    "data": [
        {
            "val1": "1125523841434490335",
            "val2": 1000025,
            "val3": "2022-03-30 14:34:57",
            "val4": 0,
            "val5": 1,
            "val6": null,
            "val7": 99.99
        }
    ],
    "rows": 1,
    "rows_before_limit_at_least": 1,
    "statistics": {
        "elapsed": 0.013290649,
        "rows_read": 61572,
        "bytes_read": 6521888
    }
}`))

	err := res.Decode()

	assert.Nil(t, err)
	assert.Empty(t, res.Documentation)
	assert.Empty(t, res.Error)
	assert.Equal(t, res.Rows, uint(1))
	assert.Equal(t, res.RowsBeforeLimitAtLeast, uint(1))
	assert.Len(t, res.Meta, 7)
	assert.NotEmpty(t, res.Statistics)
}

func TestResponseNDJSON(t *testing.T) {
	res := tinybird.Response{}
	res.Format = tinybird.NDJSON
	res.Raw = io.NopCloser(
		strings.NewReader(`{"ulid":"01H3HT0D3QG3CQRMH1SB0KKPXT","value1":12345,"value2":true,"value3":12.34}
{"ulid":"01H3HT1JB0B12QTGKH2K599B5K","value1":6543,"value2":null,"value3":99.112}`),
	)

	err := res.Decode()

	assert.Nil(t, err)
	assert.Equal(t, res.Rows, uint(2))
	assert.Len(t, res.Data, 2)
}

func TestBodyError(t *testing.T) {
	res := tinybird.Response{}
	res.Raw = io.NopCloser(strings.NewReader(`{
    "documentation": "https://docs.tinybird.co/api-reference/pipe-api.html",
    "error": "[Error] Missing columns: 'a' while processing query: 'SELECT a, b, c FROM test'"
}`))

	err := res.Decode()

	assert.Nil(t, err)
	assert.NotEmpty(t, res.Error)
	assert.NotEmpty(t, res.Documentation)
	assert.Empty(t, res.Rows)
	assert.Empty(t, res.Meta)
	assert.Empty(t, res.Data)
}

func TestBodyMalformedJSON(t *testing.T) {
	res := tinybird.Response{}
	res.Raw = io.NopCloser(strings.NewReader(`{"error: ""}`))

	err := res.Decode()

	assert.NotEmpty(t, err)
}

func TestRawIsEmpty(t *testing.T) {
	res := tinybird.Response{}

	err := res.Decode()

	assert.NotNil(t, err)
	assert.Equal(t, "Raw is empty", err.Error())
}

func TestErrorOnDecodeNDJSON(t *testing.T) {
	res := tinybird.Response{}
	res.Format = tinybird.NDJSON
	res.Raw = io.NopCloser(strings.NewReader(`{no-ndjson :(`))

	err := res.Decode()

	assert.NotNil(t, err)
}

func TestResponseHeader(t *testing.T) {
	// Create a mock HTTP response
	mockHeader := http.Header{
		"Content-Type":    {"application/json"},
		"X-Custom-Header": {"CustomValue"},
	}
	mockResponse := &http.Response{
		Header: mockHeader,
	}

	// Create a Response object and assign the mock headers
	response := tinybird.Response{}
	response.Header = mockResponse.Header

	// Validate Headers using testify assertions
	assert.Equal(t, len(mockHeader), len(response.Header), "Header count should match")

	for key, values := range mockHeader {
		assert.ElementsMatch(t, values, response.Header.Values(key), "Header values should match for key: %s", key)
	}
}

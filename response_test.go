package tinybird_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/the-hotels-network/go-tinybird"

	"github.com/stretchr/testify/assert"
)

func TestJSON(t *testing.T) {
	res := tinybird.Response{}
	res.Body = io.NopCloser(bytes.NewReader([]byte(`{
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
}`)))

	err := res.Decode()

	assert.Empty(t, res.Documentation)
	assert.Empty(t, res.Error)
	assert.Equal(t, res.Rows, uint(1))
	assert.Equal(t, res.RowsBeforeLimitAtLeast, uint(1))
	assert.Len(t, res.Meta, 7)
	assert.Nil(t, err)
	assert.NotEmpty(t, res.Statistics)
}

func TestBodyError(t *testing.T) {
	res := tinybird.Response{}
	res.Body = io.NopCloser(bytes.NewReader([]byte(`{
    "documentation": "https://docs.tinybird.co/api-reference/pipe-api.html",
    "error": "[Error] Missing columns: 'a' while processing query: 'SELECT a, b, c FROM test'"
}`)))

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
	res.Body = io.NopCloser(bytes.NewReader([]byte(`{"error: ""}`)))

	err := res.Decode()

	assert.NotEmpty(t, err)
}

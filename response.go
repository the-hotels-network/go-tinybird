package tinybird

import (
	"bytes"
	"encoding/json"

	"github.com/olivere/ndjson"
)

// Basic JSON struct for all response by Tinybird.
type Response struct {
	Body                   []byte
	Data                   []Row      `json:"data"`
	Meta                   []Meta     `json:"meta"`
	Rows                   uint       `json:"rows"`
	RowsBeforeLimitAtLeast uint       `json:"rows_before_limit_at_least"`
	Statistics             Statistics `json:"statistics"`
	Status                 int
}

type Row map[string]interface{}

// Specific field with data type.
type Meta struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// Statistics information about the response.
type Statistics struct {
	Elapsed   float64 `json:"elapsed"`
	RowsRead  uint64  `json:"rows_read"`
	BytesRead uint64  `json:"bytes_read"`
}

// Convert response body to struct.
func (r *Response) Decode() error {
	if Format() == "ndjson" {
		return r.ndjson()
	}

	return r.json()
}

// Convert body to NDJSON.
func (r *Response) ndjson() error {
	var count uint
	var rows []Row

	ndjsonReader := ndjson.NewReader(bytes.NewReader(r.Body))
	for ndjsonReader.Next() {
		var row Row
		if err := ndjsonReader.Decode(&row); err != nil {
			return err
		}
		rows = append(rows, row)
		count++
	}

	r.Data = rows
	r.Rows = count

	return nil
}

// Convert body to JSON.
func (r *Response) json() error {
	return json.Unmarshal(r.Body, &r)
}

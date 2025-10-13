package tinybird

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/olivere/ndjson"
)

// Basic JSON struct for all response by Tinybird.
type Response struct {
	Cached                 bool          // Flag control to use to determine the origin this response.
	Raw                    io.ReadCloser // Raw is the original data from response.
	Body                   string        // Body in string format, same a Data but in struct format and Raw in original format.
	Data                   Data          `json:"data,omitempty"`                       // Data is part a tinybird response.
	Documentation          string        `json:"documentation,omitempty"`              // Documentation is part a tinybird response.
	Error                  string        `json:"error,omitempty"`                      // Error is part a tinybird response.
	Meta                   []Meta        `json:"meta,omitempty"`                       // Meta is part a tinybird response.
	Rows                   uint          `json:"rows,omitempty"`                       // Rows is part a tinybird response.
	RowsBeforeLimitAtLeast uint          `json:"rows_before_limit_at_least,omitempty"` // RowsBeforeLimitAtLeast is part a tinybird response.
	Statistics             Statistics    `json:"statistics,omitempty"`                 // Statistics is part a tinybird response.
	Status                 int           // Status is a HTTP status code, ej: 200, 400, 500 etc...
	Header                 http.Header   // Headers contains the HTTP response headers returned by the server
	Format                 string        // Save format setting.
	SuccessfulRows         int           `json:"successful_rows"`
	QuarantinedRows        int           `json:"quarantined_rows"`
}

// Specific field with data type.
type Meta struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

// Statistics information about the response.
type Statistics struct {
	Elapsed   float64 `json:"elapsed,omitempty"`
	RowsRead  uint64  `json:"rows_read,omitempty"`
	BytesRead uint64  `json:"bytes_read,omitempty"`
}

// Convert response body to struct.
func (r *Response) Decode() (err error) {
	switch r.Format {
	case CSV:
		err = r.CSV()
	case NDJSON:
		err = r.NDJSON()
	default:
		err = r.JSON()
	}

	return err
}

// Convert body to CSV.
func (r *Response) CSV() error {
	if r.Raw == nil {
		return errors.New("Raw is empty")
	}

	body, err := io.ReadAll(r.Raw)
	if err != nil {
		return err
	}
	r.Body = string(body)

	return nil
}

// Convert body to NDJSON.
func (r *Response) NDJSON() error {
	ndjsonReader := ndjson.NewReader(r.Raw)
	for ndjsonReader.Next() {
		var row Row
		if err := ndjsonReader.Decode(&row); err != nil {
			return err
		}

		r.Data = append(r.Data, row)
		r.Rows++
	}

	return nil
}

// Convert body to JSON.
func (r *Response) JSON() error {
	if r.Raw == nil {
		return errors.New("Raw is empty")
	}

	body, err := io.ReadAll(r.Raw)
	if err != nil {
		return err
	}
	r.Body = string(body)

	return json.Unmarshal(body, &r)
}

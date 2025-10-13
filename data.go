package tinybird

import "encoding/json"

// Row is a generic map-based structure that can hold fields of any type.
// Each key corresponds to a field name, and each value can be any Go type.
type Row map[string]any

// Data represents a collection (slice) of rows returned from a query or dataset.
type Data []Row

// First returns the first Row in the Data slice.
// It assumes the slice is not empty.
func (d Data) First() (out Row) {
	out = d[0]

	return out
}

// FetchOne retrieves the value of a specific field from the first Row
// in the Data slice, using the provided field name (key).
// If the key doesn't exist, it returns nil.
func (d Data) FetchOne(in string) (out any) {
	row := d.First()
	out = row[in]

	return out
}

// ToString converts the entire Data slice into a JSON-formatted string.
// If marshalling fails, it returns an empty string.
func (d Data) ToString() (out string) {
	tmp, _ := json.Marshal(d)

	return string(tmp)
}

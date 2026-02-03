package tinybird

import (
	"encoding/json"
	"errors"
	"strings"
)

// Row is a generic map-based structure that can hold fields of any type.
// Each key corresponds to a field name, and each value can be any Go type.
type Row map[string]any

// Data represents a collection (slice) of rows returned from a query or dataset.
type Data []Row

// Count the items.
func (d Data) Len() int {
	return len(d)
}

// First returns the first Row in the Data slice.
// It assumes the slice is not empty.
func (d Data) First() Row {
	if d.Len() > 0 {
		return d[0]
	}

	return Row{}
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

// Get retrieves a value from the first row of the Data structure
// by traversing a dot-separated path.
//
// The path represents nested keys within Row maps (e.g. "a.b.c").
// If the Data slice is empty, the path does not exist, or an
// intermediate value is not a Row, Get returns nil.
func (d Data) Get(in string) any {
	if len(d) == 0 {
		return nil
	}

	parts := strings.Split(in, ".")

	for _, row := range d {
		if v, ok := get(row, parts); ok {
			return v
		}
	}

	return nil
}

func get(row Row, parts []string) (any, bool) {
	var current any = row

	for _, part := range parts {
		r, ok := current.(Row)
		if !ok {
			return nil, false
		}

		current, ok = r[part]
		if !ok {
			return nil, false
		}
	}

	return current, true
}

// Set assigns a value in the first row of the Data structure
// at the location specified by a dot-separated path.
//
// Intermediate Row nodes are created as needed if they do not exist.
// An error is returned if Data is empty, the path is empty, or
// a non-Row value is encountered while traversing the path.
func (d *Data) Set(in string, value any) error {
	if d == nil {
		return errors.New("nil Data")
	}

	parts := strings.Split(in, ".")

	for i := range *d {
		row := (*d)[i]

		if _, ok := get(row, parts[:len(parts)-1]); ok {
			return set(row, parts, value)
		}
	}

	return errors.New("path not found")
}

func set(row Row, parts []string, value any) error {
	current := row

	for i, p := range parts {
		if i == len(parts)-1 {
			current[p] = value
			return nil
		}

		next, exists := current[p]
		if !exists {
			child := Row{}
			current[p] = child
			current = child
			continue
		}

		child, ok := next.(Row)
		if !ok {
			return errors.New("path collision at " + p)
		}

		current = child
	}

	return nil
}

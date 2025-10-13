package tinybird_test

import (
	"testing"

	"github.com/the-hotels-network/go-tinybird"

	"github.com/stretchr/testify/assert"
)

func TestData_Empty(t *testing.T) {
	d1 := tinybird.Data{}

	assert.Equal(t, tinybird.Row{}, d1.First())
	assert.Equal(t, 0, d1.Count())
	assert.Equal(t, nil, d1.FetchOne("foo"))
	assert.Equal(t, "[]", d1.ToString())

	var d2 tinybird.Data

	assert.Equal(t, tinybird.Row{}, d2.First())
	assert.Equal(t, 0, d2.Count())
	assert.Equal(t, nil, d2.FetchOne("foo"))
	assert.Equal(t, "[]", d2.ToString())
}

func TestData_First(t *testing.T) {
	d := tinybird.Data{
		{"id": 1, "name": "first"},
		{"id": 2, "name": "second"},
	}

	assert.Equal(t, 2, d.Count())
	assert.Equal(t, tinybird.Row{"id": 1, "name": "first"}, d.First())
}

func TestData_FetchOne(t *testing.T) {
	d := tinybird.Data{
		{"id": 10, "name": "first", "bool": true},
	}

	assert.Equal(t, 1, d.Count())
	assert.Equal(t, "first", d.FetchOne("name"))
	assert.Equal(t, true, d.FetchOne("bool"))
	assert.Equal(t, 10, d.FetchOne("id"))
	assert.Nil(t, d.FetchOne("missing"))
}

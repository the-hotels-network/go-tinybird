package tinybird_test

import (
	"testing"

	"github.com/the-hotels-network/go-tinybird"

	"github.com/stretchr/testify/assert"
)

func TestFirst_ReturnsFirstRow(t *testing.T) {
	d := tinybird.Data{
		{"id": 1, "name": "first"},
		{"id": 2, "name": "second"},
	}

	first := d.First()
	assert.Equal(t, tinybird.Row{"id": 1, "name": "first"}, first)
}

func TestFirst_PanicsOnEmptyData(t *testing.T) {
	var d tinybird.Data
	assert.Equal(t, d.First(), tinybird.Row{})

	empty := tinybird.Data{}
	assert.Equal(t, empty.First(), tinybird.Row{})
}

func TestFetchOne_ReturnsValueFromFirstRow(t *testing.T) {
	d := tinybird.Data{
		{"foo": "bar", "n": 42},
		{"foo": "baz"},
	}

	val := d.FetchOne("foo")
	assert.Equal(t, "bar", val)

	valNum := d.FetchOne("n")
	assert.Equal(t, 42, valNum)
}

func TestFetchOne_ReturnsNilWhenKeyMissing(t *testing.T) {
	d := tinybird.Data{
		{"foo": "bar"},
	}

	val := d.FetchOne("missing")
	assert.Nil(t, val)
}

func TestFetchOne_PanicsOnEmptyData(t *testing.T) {
	var d tinybird.Data
	assert.Nil(t, d.FetchOne("foo"))
}

func TestToString_JSONRoundTripEq(t *testing.T) {
	d := tinybird.Data{
		{"a": 1, "b": "x"},
		{"a": 2, "c": true},
	}

	expected := `[{"a":1,"b":"x"},{"a":2,"c":true}]`
	assert.JSONEq(t, expected, d.ToString())
}

func TestToString_OnNilAndEmpty(t *testing.T) {
	var dNil tinybird.Data
	dEmpty := tinybird.Data{}

	assert.Equal(t, "null", dNil.ToString())
	assert.Equal(t, "[]", dEmpty.ToString())
}

func TestGet(t *testing.T) {
	d := tinybird.Data{
		{"a": nil},
		{"b": tinybird.Row{"a": 1}},
		{"b": tinybird.Row{"b": 2}},
		{"b": tinybird.Row{"c": 3}},
		{"c": tinybird.Row{"a": 4}},
		{"d": tinybird.Row{"a": 5}},
		{"d": tinybird.Row{"b": tinybird.Row{"a": 6}}},
		{"f": 7},
	}

	assert.Nil(t, d.Get(""))
	assert.Nil(t, d.Get("a"))
	assert.Nil(t, d.Get("e"))
	assert.Equal(t, 7, d.Get("f"))
	assert.Equal(t, 4, d.Get("c.a"))
	assert.Equal(t, 6, d.Get("d.b.a"))
	assert.Equal(t, tinybird.Row{"a": 4}, d.Get("c"))
}

func TestSet(t *testing.T) {
	d := tinybird.Data{
		{"a": nil},
		{"b": tinybird.Row{"a": 1}},
		{"b": tinybird.Row{"b": 2}},
		{"b": tinybird.Row{"c": 3}},
		{"c": tinybird.Row{"a": 4}},
		{"d": tinybird.Row{"a": 5}},
		{"d": tinybird.Row{"b": tinybird.Row{"a": 6}}},
		{"f": 7},
	}

	assert.Nil(t, d.Set("a", 0))
	assert.Equal(t, 0, d.Get("a"))
	assert.Nil(t, d.Set("b.b", 22))
	assert.Equal(t, 22, d.Get("b.b"))
	assert.Nil(t, d.Set("d.b.a", 66))
	assert.Equal(t, 66, d.Get("d.b.a"))
}

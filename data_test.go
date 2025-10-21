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

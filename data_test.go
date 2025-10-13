package tinybird

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirst_ReturnsFirstRow(t *testing.T) {
	d := Data{
		{"id": 1, "name": "first"},
		{"id": 2, "name": "second"},
	}

	first := d.First()
	assert.Equal(t, Row{"id": 1, "name": "first"}, first)
}

func TestFirst_PanicsOnEmptyData(t *testing.T) {
	var d Data
	assert.Panics(t, func() { _ = d.First() })

	empty := Data{}
	assert.Panics(t, func() { _ = empty.First() })
}

func TestFetchOne_ReturnsValueFromFirstRow(t *testing.T) {
	d := Data{
		{"foo": "bar", "n": 42},
		{"foo": "baz"},
	}

	val := d.FetchOne("foo")
	assert.Equal(t, "bar", val)

	valNum := d.FetchOne("n")
	assert.Equal(t, 42, valNum)
}

func TestFetchOne_ReturnsNilWhenKeyMissing(t *testing.T) {
	d := Data{
		{"foo": "bar"},
	}

	val := d.FetchOne("missing")
	assert.Nil(t, val)
}

func TestFetchOne_PanicsOnEmptyData(t *testing.T) {
	var d Data
	assert.Panics(t, func() { _ = d.FetchOne("foo") })
}

func TestToString_JSONRoundTripEq(t *testing.T) {
	d := Data{
		{"a": 1, "b": "x"},
		{"a": 2, "c": true},
	}

	expected := `[{"a":1,"b":"x"},{"a":2,"c":true}]`
	assert.JSONEq(t, expected, d.ToString())
}

func TestToString_OnNilAndEmpty(t *testing.T) {
	var dNil Data
	dEmpty := Data{}

	assert.Equal(t, "null", dNil.ToString())
	assert.Equal(t, "[]", dEmpty.ToString())
}

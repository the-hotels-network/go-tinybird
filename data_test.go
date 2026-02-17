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

func TestAutoCast(t *testing.T) {
	tests := []struct {
		name     string
		val      any
		expected any
	}{
		// Non-string values pass through unchanged
		{"nil", nil, nil},
		{"int", 42, 42},
		{"float64", float64(3.14), float64(3.14)},
		{"bool", true, true},

		// Strings that are integers
		{"string int", "42", int64(42)},
		{"string negative int", "-999", int64(-999)},
		{"string big int", "1125523841434490335", int64(1125523841434490335)},
		{"string zero", "0", int64(0)},

		// Strings that are floats
		{"string float", "3.14", float64(3.14)},
		{"string negative float", "-0.5", float64(-0.5)},
		{"string float with exponent", "1.5e2", float64(150)},

		// Strings that stay as strings
		{"plain string", "hello", "hello"},
		{"empty string", "", ""},
		{"date string", "2022-03-30", "2022-03-30"},
		{"datetime string", "2022-03-30 14:34:57", "2022-03-30 14:34:57"},
		{"mixed string", "abc123", "abc123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tinybird.AutoCast(tt.val)
			assert.Equal(t, tt.expected, got)
		})
	}
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

func TestSetWithAutoCast(t *testing.T) {
	d := tinybird.Data{
		{"a": nil},
	}

	assert.Nil(t, d.SetWithAutoCast("a", "42"))
	assert.Equal(t, int64(42), d.Get("a"))

	assert.Nil(t, d.SetWithAutoCast("a", "3.14"))
	assert.Equal(t, float64(3.14), d.Get("a"))

	assert.Nil(t, d.SetWithAutoCast("a", "hello"))
	assert.Equal(t, "hello", d.Get("a"))

	assert.Nil(t, d.SetWithAutoCast("a", 99))
	assert.Equal(t, 99, d.Get("a"))
}

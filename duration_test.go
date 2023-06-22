package tinybird_test

import (
	"errors"
	"testing"
	"time"

	"github.com/the-hotels-network/go-tinybird"

	"github.com/stretchr/testify/assert"
)

func TestDurationElapsed(t *testing.T) {
	var d tinybird.Duration
	d.Do(func() error {
		time.Sleep(1 * time.Second)
		return errors.New("Test")
	})

	assert.Equal(t, d.Seconds()[0:3], "1.0")
}

func TestDurationError(t *testing.T) {
	var d tinybird.Duration
	err := d.Do(func() error {
		return errors.New("Test")
	})

	assert.Equal(t, err, errors.New("Test"))
}

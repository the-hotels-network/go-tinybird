package tinybird

import (
	"fmt"
	"time"
)

type Duration time.Duration

// Calculate child function duration.
func (d *Duration) Do(fn func() error) error {
	t1 := time.Now()
	err := fn()

	(*d) = Duration(time.Since(t1))

	return err
}

// Convert duration to seconds unit.
func (d Duration) Seconds() string {
	return fmt.Sprintf(
		"%.2fs",
		time.Duration(d).Seconds(),
	)
}

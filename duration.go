package tinybird

import (
	"fmt"
	"time"
)

// Calculate child function duration.
func Duration(fn func() error) (string, error) {
	t1 := time.Now()
	err := fn()
	t2 := time.Now()
	diff := t2.Sub(t1)
	out := time.Time{}.Add(diff)

	return fmt.Sprint(out.Format("15:04:05.000")), err
}

package tinybird

import (
	"sync"
	"time"
)

// Slice for all request to execute in parallel.
type Requests []Request

// Add request into slice.
func (rs *Requests) Add(i Request) {
	*rs = append(*rs, i)
}

// Execute multithreading/parallel request.
func (rs *Requests) Execute() {
	wg := sync.WaitGroup{}
	for index, _ := range *rs {
		wg.Add(1)
		go func(i int) {
			(*rs)[i].Execute()
			wg.Done()
		}(index)
	}
	wg.Wait()
}

// Return elapsed time for all requests.
func (rs Requests) Duration() Duration {
	var d time.Duration

	for _, r := range rs {
		d += time.Duration(r.Elapsed)
	}
	return Duration(d)
}

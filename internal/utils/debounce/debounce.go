// This file was heavily inspired by `https://github.com/romdo/go-debounce`
package debounce

import (
	"sync"
	"time"
)

type Debouncer struct {
	Debounced func() bool
	Cancelled func() bool
	Set       func(args ...any)
}

func New(wait time.Duration, f func(args ...any)) *Debouncer {
	var mu sync.Mutex
	var latest []any

	timer := time.AfterFunc(24*time.Hour, func() {
		mu.Lock()
		defer mu.Unlock()
		f(latest...)
	})
	timer.Stop()

	debounced := func() bool {
		mu.Lock()
		defer mu.Unlock()
		return timer.Reset(wait)
	}

	cancelled := func() bool {
		mu.Lock()
		defer mu.Unlock()
		return timer.Stop()
	}

	set := func(args ...any) {
		mu.Lock()
		defer mu.Unlock()
		latest = args
	}

	return &Debouncer{
		Debounced: debounced,
		Cancelled: cancelled,
		Set:       set,
	}
}

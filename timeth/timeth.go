// Package timeth provide a hal.Time implementation, for easier to set/tick
// current mock time.
// hal.Time is auto replaced(mocked) at the first call of Set/Tick() function,
// and auto reset in testing/reset. Can also manual set by calling Install()
// function, if auto install cause data-race condition.
//
// Time default to 2016-1-1 local timezone.
package timeth

import (
	"spork/testing/reset"
	"sync"
	"time"

	"github.com/redforks/hal"
)

var (
	l       = sync.Mutex{}
	mocked  = false
	current time.Time
)

func init() {
	reset.Register(nil, func() {
		l.Lock()
		current = time.Date(2016, 1, 1, 0, 0, 0, 0, time.Local)
		hal.Now = time.Now
		mocked = false
		l.Unlock()
	})
}

// Install mocked version of hal.Now.
func Install() {
	l.Lock()
	install()
	l.Unlock()
}

// Set mocked time to t, auto install time mock if not mocked.
func Set(t time.Time) {
	l.Lock()
	install()
	current = t
	l.Unlock()
}

// Tick current time by d duration, auto install time mock if not mocked.
func Tick(d time.Duration) {
	l.Lock()
	install()
	current = current.Add(d)
	l.Unlock()
}

func install() {
	if !mocked {
		// We can not re-assign hal.Now every time, it cause data-race condition,
		// when other goroutie calling hal.Now at the same time.
		hal.Now = now
	}
}

// Mocked version of hal.Now
func now() time.Time {
	return current
}

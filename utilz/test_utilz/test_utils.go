// +build testing

package test_utilz

import (
	"github.com/explodes/serving/utilz"
	"time"
)

var _ utilz.Clock = (*TestClock)(nil)

type TestClock struct {
	t time.Time
}

// NewTestClock creates a clock used for testing starting t=0.
func NewTestClock() *TestClock {
	return &TestClock{
		t: time.Time{},
	}
}

// Now returns the current time of the clock.
func (c *TestClock) Now() time.Time {
	return c.t
}

// Add increases the time of the clock.
func (c *TestClock) Add(d time.Duration) time.Time {
	c.t = c.t.Add(d)
	return c.t
}

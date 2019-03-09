// +build testing

package test_utilz

import (
	spb "github.com/explodes/serving/proto"
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
		t: time.Unix(0, 0),
	}
}

// Now gets the current time on the clock.
func (c *TestClock) Now() time.Time {
	return c.t
}

// Timestamp gets a Timestamp of the current time.
func (c *TestClock) Timestamp() *spb.Timestamp {
	return spb.TimestampTime(c.Now())
}

// Add increases the time of the clock.
func (c *TestClock) Add(d time.Duration) time.Time {
	c.t = c.t.Add(d)
	return c.t
}

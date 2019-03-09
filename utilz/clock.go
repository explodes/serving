package utilz

import (
	spb "github.com/explodes/serving/proto"
	"time"
)

// Clock gets the current time.
type Clock interface {
	// Now gets the current time on the clock.
	Now() time.Time
	// Timestamp gets a Timestamp of the current time.
	Timestamp() *spb.Timestamp
}

var _ Clock = realClock{}

// realClock implements Clock but uses the real system time.
type realClock struct{}

// NewClock creates a Clock based on system time.
func NewClock() Clock {
	return realClock{}
}

// Now gets the current time on the clock.
func (r realClock) Now() time.Time {
	return time.Now()
}

// Timestamp gets a Timestamp of the current time.
func (r realClock) Timestamp() *spb.Timestamp {
	return spb.TimestampTime(r.Now())
}

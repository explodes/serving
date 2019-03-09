package utilz

import "time"

// Clock gets the current time.
type Clock interface {
	// Now gets the current time on the clock.
	Now() time.Time
}

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

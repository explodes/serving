package proto

import "time"

func DurationNanos(nanos int64) *Duration {
	return &Duration{
		Unit: &Duration_Nanoseconds{Nanoseconds: nanos},
	}
}

func (m *Duration) Duration() time.Duration {
	switch t := m.Unit.(type) {
	case *Duration_Nanoseconds:
		return time.Duration(t.Nanoseconds) * time.Nanosecond
	case *Duration_Milliseconds:
		return time.Duration(t.Milliseconds) * time.Millisecond
	case *Duration_Seconds:
		return time.Duration(t.Seconds) * time.Second
	}
	return 0
}

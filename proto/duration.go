package proto

import "time"

func DurationNanos(nanos int64) *Duration {
	return &Duration{
		Unit: &Duration_Nanoseconds{Nanoseconds: nanos},
	}
}

func DurationMillis(millis int64) *Duration {
	return &Duration{
		Unit: &Duration_Milliseconds{Milliseconds: millis},
	}
}

func DurationSeconds(seconds int64) *Duration {
	return &Duration{
		Unit: &Duration_Seconds{Seconds: seconds},
	}
}

func (m *Duration) Duration() time.Duration {
	if m.Unit == nil {
		return 0
	}
	switch t := m.Unit.(type) {
	case *Duration_Nanoseconds:
		return time.Duration(t.Nanoseconds) * time.Nanosecond
	case *Duration_Milliseconds:
		return time.Duration(t.Milliseconds) * time.Millisecond
	case *Duration_Seconds:
		return time.Duration(t.Seconds) * time.Second
	default:
		panic("unknown duration type")
	}
}

package proto

import "time"

func TimestampNow() *Timestamp {
	return TimestampTime(time.Now())
}

func TimestampTime(t time.Time) *Timestamp {
	return &Timestamp{
		Unit: &Timestamp_Nanoseconds{
			Nanoseconds: t.UnixNano(),
		},
	}
}

func (m *Timestamp) Time() time.Time {
	if m.Unit == nil {
		return time.Unix(0, 0)
	}
	switch t := m.Unit.(type) {
	case *Timestamp_Nanoseconds:
		return time.Unix(0, t.Nanoseconds)
	case *Timestamp_Milliseconds:
		return time.Unix(0, t.Milliseconds*int64(time.Millisecond))
	case *Timestamp_Seconds:
		return time.Unix(t.Seconds, 0)
	default:
		panic("unknown time unit")
	}
}

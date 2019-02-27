package proto

import "time"

func Now() *Timestamp {
	now := time.Now()
	return &Timestamp{
		Unit: &Timestamp_Nanoseconds{
			Nanoseconds: now.UnixNano(),
		},
	}
}

func (m *Timestamp) Time() time.Time {
	switch t := m.Unit.(type) {
	case *Timestamp_Nanoseconds:
		return time.Unix(t.Nanoseconds/1e9, t.Nanoseconds%1e9)
	case *Timestamp_Milliseconds:
		return time.Unix(t.Milliseconds/1e3, t.Milliseconds%1e3)
	case *Timestamp_Seconds:
		return time.Unix(t.Seconds, 0)
	}
	return time.Unix(0, 0)
}

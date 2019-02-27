package serving

import "time"

func Now() *Timestamp {
	now := time.Now()
	return &Timestamp{
		Milliseconds: int64(time.Duration(now.UnixNano()) / time.Millisecond),
	}
}

func (t *Timestamp) Time() time.Time {
	nanos := int64(time.Duration(t.Milliseconds) * time.Millisecond)
	return time.Unix(nanos/1e9, nanos%1e9)
}
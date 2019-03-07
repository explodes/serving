package proto

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimestampNow(t *testing.T) {
	ts := TimestampNow()
	now := time.Now()
	tsNanos := ts.Time().UnixNano()
	assert.True(t, tsNanos > now.UnixNano()-1000)
}

func TestTimestampTime(t *testing.T) {
	when := time.Now().Add(-time.Hour * 5)
	ts := TimestampTime(when)
	assert.Equal(t, when.UnixNano(), ts.Time().UnixNano())
}

func TestTimestamp_Time_Nanos(t *testing.T) {
	ts := &Timestamp{Unit: &Timestamp_Nanoseconds{Nanoseconds: 100}}
	nanos := ts.Time().UnixNano()
	assert.Equal(t, int64(100*time.Nanosecond), nanos)
}

func TestTimestamp_Time_Millis(t *testing.T) {
	ts := &Timestamp{Unit: &Timestamp_Milliseconds{Milliseconds: 100}}
	nanos := ts.Time().UnixNano()
	assert.Equal(t, int64(100*time.Millisecond), nanos)
}

func TestTimestamp_Time_Seconds(t *testing.T) {
	ts := &Timestamp{Unit: &Timestamp_Seconds{Seconds: 100}}
	nanos := ts.Time().UnixNano()
	assert.Equal(t, int64(100*time.Second), nanos)
}


func TestTimestamp_Time_Nil(t *testing.T) {
	ts := &Timestamp{Unit: nil}
	nanos := ts.Time().UnixNano()
	assert.Equal(t, int64(0), nanos)
}

type unknownTime struct{}

func (u *unknownTime) isTimestamp_Unit() {}

func TestTimestamp_Time_unknown(t *testing.T) {
	defer func() {
		err := recover()
		assert.NotNil(t, err)
	}()
	pb := &Timestamp{Unit: &unknownTime{}}
	pb.Time()
}

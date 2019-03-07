package proto

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDurationNanos(t *testing.T) {
	pb := DurationNanos(100)
	d := pb.Duration()

	assert.Equal(t, time.Duration(100)*time.Nanosecond, d)
}

func TestDurationMillis(t *testing.T) {
	pb := DurationMillis(100)
	d := pb.Duration()

	assert.Equal(t, time.Duration(100)*time.Millisecond, d)
}

func TestDurationSeconds(t *testing.T) {
	pb := DurationSeconds(100)
	d := pb.Duration()

	assert.Equal(t, time.Duration(100)*time.Second, d)
}

func TestDuration_Duration_nil(t *testing.T) {
	pb := &Duration{Unit: nil}
	d := pb.Duration()
	assert.Equal(t, 0*time.Second, d)
}

type unknownDuration struct{}

func (u *unknownDuration) isDuration_Unit() {}

func TestDuration_Duration_unknown(t *testing.T) {
	defer func() {
		err := recover()
		assert.NotNil(t, err)
	}()
	pb := &Duration{Unit: &unknownDuration{}}
	pb.Duration()
}

package statusz

import (
	"github.com/explodes/serving/utilz/test_utilz"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewRateTracker(t *testing.T) {
	rt := NewRateTracker("test")
	assert.NotNil(t, rt)
}

func TestNewRateTrackerClock(t *testing.T) {
	clock := test_utilz.NewTestClock()
	rt := NewRateTrackerClock("test", clock)
	assert.NotNil(t, rt)

	log := rt.Start()
	clock.Add(1 * time.Hour)
	log.End()

	metrics, err := rt.MarshalMetrics()
	assert.NoError(t, err)
	assert.Len(t, metrics, 2)

	assert.Equal(t, metrics[0].Name, "test.count")
	assert.Equal(t, uint64(1), metrics[0].Value.(*Metric_U64).U64)

	assert.Equal(t, metrics[1].Name, "test.avg_duration")
	assert.Equal(t, int64(1*time.Hour), metrics[1].Value.(*Metric_Duration).Duration.GetNanoseconds())
}

func TestNewRateTrackerClock_zero(t *testing.T) {
	clock := test_utilz.NewTestClock()
	rt := NewRateTrackerClock("test", clock)
	assert.NotNil(t, rt)

	metrics, err := rt.MarshalMetrics()
	assert.NoError(t, err)
	assert.Len(t, metrics, 2)

	assert.Equal(t, metrics[0].Name, "test.count")
	assert.Equal(t, uint64(0), metrics[0].Value.(*Metric_U64).U64)

	assert.Equal(t, metrics[1].Name, "test.avg_duration")
	assert.Equal(t, int64(0), metrics[1].Value.(*Metric_Duration).Duration.GetNanoseconds())
}

func TestNewRateTrackerClock_limit(t *testing.T) {
	clock := test_utilz.NewTestClock()
	rt := NewRateTrackerClock("test", clock)
	assert.NotNil(t, rt)

	for i := 0; i < rateHistorySize; i++ {
		log := rt.Start()
		clock.Add(1 * time.Minute)
		log.End()
	}
	for i := 0; i < rateHistorySize; i++ {
		log := rt.Start()
		clock.Add(1 * time.Hour)
		log.End()
	}

	metrics, err := rt.MarshalMetrics()
	assert.NoError(t, err)
	assert.Len(t, metrics, 2)

	assert.Equal(t, metrics[0].Name, "test.count")
	assert.Equal(t, uint64(2*rateHistorySize), metrics[0].Value.(*Metric_U64).U64)

	assert.Equal(t, metrics[1].Name, "test.avg_duration")
	assert.Equal(t, int64(1*time.Hour), metrics[1].Value.(*Metric_Duration).Duration.GetNanoseconds())
}

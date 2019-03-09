package statusz_test

import (
	"github.com/explodes/serving/statusz"
	"github.com/explodes/serving/test_serving"
	"github.com/explodes/serving/utilz/test_utilz"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStatuszServer_GetStatus(t *testing.T) {
	clock := test_utilz.NewTestClock()
	clock.Add(1 * time.Nanosecond)
	s := statusz.NewStatuszServerClock(clock)

	req := &statusz.GetStatusRequest{Cookie: "somecookie"}
	res, err := s.GetStatus(test_serving.TestContext(), req)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	status := res.Status
	assert.Equal(t, int64(1), status.Timestamp.GetNanoseconds())
	assert.Len(t, status.Groups, 2)
	assert.Equal(t, "system", status.Groups[0].Name)
	assert.Len(t, status.Groups[0].Metrics, 3)
	group := status.Groups[1]
	assert.Equal(t, "Statusz", group.Name)
	assert.Len(t, group.Metrics, 2)
	assert.Equal(t, "GetStatus.count", group.Metrics[0].Name)
	assert.Equal(t, uint64(0), group.Metrics[0].Value.(*statusz.Metric_U64).U64)
	assert.Equal(t, group.Metrics[1].Name, "GetStatus.avg_duration")
	assert.Equal(t, 0*time.Nanosecond, group.Metrics[1].Value.(*statusz.Metric_Duration).Duration.Duration())

	// The recording of a request status is deferred.
	// We need to call this a second time to get
	// the Vars calculated the from the first request.
	clock.Add(1 * time.Nanosecond)
	res, err = s.GetStatus(test_serving.TestContext(), req)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	status = res.Status
	assert.Equal(t, int64(2), status.Timestamp.GetNanoseconds())
	assert.Len(t, status.Groups, 2)
	assert.Equal(t, "system", status.Groups[0].Name)
	assert.Len(t, status.Groups[0].Metrics, 3)
	group = status.Groups[1]
	assert.Equal(t, "Statusz", group.Name)
	assert.Len(t, group.Metrics, 2)
	assert.Equal(t, "GetStatus.count", group.Metrics[0].Name)
	assert.Equal(t, uint64(1), group.Metrics[0].Value.(*statusz.Metric_U64).U64)
	assert.Equal(t, group.Metrics[1].Name, "GetStatus.avg_duration")
	assert.Equal(t, 0*time.Nanosecond, group.Metrics[1].Value.(*statusz.Metric_Duration).Duration.Duration())
}

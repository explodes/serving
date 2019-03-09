package statusz

import (
	"fmt"
	spb "github.com/explodes/serving/proto"
	"github.com/explodes/serving/utilz"
	"sync"
)

const (
	RateHistorySize = 1000
)

type Ender interface {
	End()
}

var _ Var = (*RateTracker)(nil)

type RateTracker struct {
	mu         *sync.RWMutex
	clock      utilz.Clock
	name       string
	count      uint64
	timeframes *timeframeRing
}

func NewRateTracker(name string) *RateTracker {
	return NewRateTrackerClock(name, utilz.NewClock())
}

func NewRateTrackerClock(name string, clock utilz.Clock) *RateTracker {
	return &RateTracker{
		mu:         &sync.RWMutex{},
		clock:      clock,
		name:       name,
		count:      0,
		timeframes: newTimeframeRing(),
	}

}

func (r *RateTracker) MarshalMetrics() ([]*Metric, error) {
	count, avgDuration := r.stats()
	return []*Metric{
		{Name: fmt.Sprintf("%s.count", r.name), Value: &Metric_U64{U64: count}},
		{Name: fmt.Sprintf("%s.avg_duration", r.name), Value: &Metric_Duration{Duration: spb.DurationNanos(int64(avgDuration))}},
	}, nil
}

func (r *RateTracker) Start() Ender {
	return &deferredRateLogger{
		r:     r,
		start: r.clock.Now().UnixNano(),
	}
}

func (r *RateTracker) record(start, end int64) {
	r.mu.Lock()
	r.count++
	r.timeframes.put(timeframe{start, end})
	r.mu.Unlock()
}

func (r *RateTracker) stats() (count uint64, avgDuration int64) {
	r.mu.RLock()

	// Calculate the count.
	count = r.count

	// Calculate the average duration of the last N timeframes.
	N := r.timeframes.len()
	if N == 0 {
		avgDuration = 0
	} else {
		sumDiff := int64(0)
		for i := 0; i < N; i++ {
			tf := r.timeframes.timeframes[i]
			sumDiff += tf.end - tf.start
		}
		avgDuration = sumDiff / int64(N)
	}

	r.mu.RUnlock()
	return
}

type timeframe struct {
	start, end int64
}

type timeframeRing struct {
	size, offset int
	timeframes   [RateHistorySize]timeframe
}

func newTimeframeRing() *timeframeRing {
	return &timeframeRing{
		size:       0,
		offset:     0,
		timeframes: [RateHistorySize]timeframe{},
	}
}

func (r *timeframeRing) len() int {
	return r.size
}

func (r *timeframeRing) put(t timeframe) {
	if r.size < RateHistorySize {
		r.size++
	}
	index := r.offset % RateHistorySize
	r.timeframes[index] = t
	r.offset = index + 1
}

var _ Ender = (*deferredRateLogger)(nil)

type deferredRateLogger struct {
	r     *RateTracker
	start int64
}

func (d *deferredRateLogger) End() {
	now := d.r.clock.Now().UnixNano()
	d.r.record(d.start, now)
}

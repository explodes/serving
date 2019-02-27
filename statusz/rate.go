package statusz

import (
	"fmt"
	spb "github.com/explodes/serving/proto"
	"sync"
	"time"
)

const (
	rateHistorySize = 1000
)

type EndTimeLogger interface {
	End()
}

var _ Var = (*RateTracker)(nil)

type RateTracker struct {
	mu         *sync.RWMutex
	name       string
	count      uint64
	timeframes *timeframeRingBuffer
}

func NewRateTracker(name string) *RateTracker {
	return &RateTracker{
		mu:         &sync.RWMutex{},
		name:       name,
		count:      0,
		timeframes: newTimeframeRingBuffer(),
	}
}

func (r *RateTracker) Marshal() ([]*Metric, error) {
	count, avgDuration := r.stats()
	return []*Metric{
		{Name: fmt.Sprintf("%s.count", r.name), Value: &Metric_U64{U64: count}},
		{Name: fmt.Sprintf("%s.avg_duration", r.name), Value: &Metric_Duration{Duration: spb.DurationNanos(int64(avgDuration))}},
	}, nil
}

func (r *RateTracker) Start() EndTimeLogger {
	return &deferredRateLogger{
		r:     r,
		start: time.Now().UnixNano(),
	}
}

func (r *RateTracker) record(start, end int64) {
	r.mu.Lock()
	r.count++
	r.timeframes.put(timeframe{start, end})
	r.mu.Unlock()
}

func (r *RateTracker) stats() (count uint64, avgDuration float64) {
	r.mu.RLock()
	count = r.count

	if r.timeframes.len() == 0 {
		avgDuration = 0
	} else {
		sumDiff := int64(0)
		for i := 0; i < r.timeframes.len(); i++ {
			tf := r.timeframes.timeframes[i]
			sumDiff += tf.end - tf.start
		}
		avgDuration = float64(sumDiff) / float64(r.timeframes.len())
	}

	r.mu.RUnlock()
	return
}

type timeframe struct {
	start, end int64
}

type timeframeRingBuffer struct {
	size, offset int
	timeframes   [rateHistorySize]timeframe
}

func newTimeframeRingBuffer() *timeframeRingBuffer {
	return &timeframeRingBuffer{
		size:       0,
		offset:     0,
		timeframes: [rateHistorySize]timeframe{},
	}
}

func (r *timeframeRingBuffer) len() int {
	return r.size
}

func (r *timeframeRingBuffer) put(t timeframe) {
	if r.size < rateHistorySize {
		r.size++
	}
	index := r.offset % rateHistorySize
	r.timeframes[index] = t
	r.offset = index + 1
}

var _ EndTimeLogger = (*deferredRateLogger)(nil)

type deferredRateLogger struct {
	r     *RateTracker
	start int64
}

func (d *deferredRateLogger) End() {
	now := time.Now().UnixNano()
	d.r.record(d.start, now)
}

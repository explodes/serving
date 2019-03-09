package statusz

import (
	"fmt"
	"github.com/explodes/serving/utilz"
	"golang.org/x/net/context"
	"sync"
)

var (
	registerOnce              = &sync.Once{}
	varGetStatus *RateTracker = nil
)

func registerServerVars(clock utilz.Clock) {
	varGetStatus = NewRateTrackerClock("GetStatus", clock)
	Register("Statusz", VarGroup{
		varGetStatus,
	})
}

type statuszServer struct {
	clock utilz.Clock
}

func NewStatuszServer() StatuszServiceServer {
	return NewStatuszServerClock(utilz.NewClock())
}

func NewStatuszServerClock(clock utilz.Clock) StatuszServiceServer {
	registerOnce.Do(func() {
		registerServerVars(clock)
	})
	return &statuszServer{
		clock: clock,
	}
}

func (s *statuszServer) GetStatus(ctx context.Context, req *GetStatusRequest) (*GetStatusResponse, error) {
	defer varGetStatus.Start().End()

	groups, err := collectMetricGroups()
	if err != nil {
		return nil, err
	}
	res := &GetStatusResponse{
		Status: &Status{
			Timestamp: s.clock.Timestamp(),
			Groups:    groups,
		},
	}
	return res, nil
}

func collectMetricGroups() ([]*MetricGroup, error) {
	groups := make([]*MetricGroup, 0, len(varRegistry))
	for _, namedVar := range varRegistry {
		name := namedVar.name
		v := namedVar.v
		metrics, err := v.MarshalMetrics()
		if err != nil {
			return nil, fmt.Errorf("error marshalling metric %s: %v", name, err)
		}
		group := &MetricGroup{
			Name:    name,
			Metrics: metrics,
		}
		groups = append(groups, group)

	}
	return groups, nil
}

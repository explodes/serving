package statusz

import (
	"fmt"
	spb "github.com/explodes/serving/proto"
	"golang.org/x/net/context"
	"sync"
)

var (
	registerOnce = &sync.Once{}
	varGetStatus = NewRateTracker("GetStatus")
)

func registerServerVars() {
	Register("Statusz", VarGroup{
		varGetStatus,
	})
}

type statuszServer struct{}

func NewStatuszServer() StatuszServiceServer {
	registerOnce.Do(registerServerVars)
	return &statuszServer{}
}

func (s *statuszServer) GetStatus(ctx context.Context, req *GetStatusRequest) (*GetStatusResponse, error) {
	defer varGetStatus.Start().End()

	groups, err := collectMetricGroups()
	if err != nil {
		return nil, err
	}
	res := &GetStatusResponse{
		Status: &Status{
			Timestamp: spb.TimestampNow(),
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

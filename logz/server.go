package logz

import (
	"context"
	"github.com/explodes/serving/statusz"
	"sync"
)

var _ LogzServiceServer = (*logzServer)(nil)

var (
	registerOnce = &sync.Once{}
	varRecord = statusz.NewRateTracker("Record")
)

func registerServerVars() {
	statusz.Register("Logz", statusz.VarGroup{
		varRecord,
	})
}

type logzServer struct {
	backend Backend
}

func NewLogzServer(config *ServiceConfig) LogzServiceServer {
	backend := NewConsoleBackend()
	return NewLogzServerBackend(backend)
}

func NewLogzServerBackend(backend Backend) LogzServiceServer {
	registerOnce.Do(registerServerVars)
	return &logzServer{
		backend: backend,
	}
}

func (s *logzServer) Record(ctx context.Context, req *RecordRequest) (*RecordResponse, error) {
	defer varRecord.Start().End()

	if err := s.backend.Record(req); err != nil {
		return nil, err
	}
	res := &RecordResponse{}
	return res, nil
}

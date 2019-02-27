package logz

import (
	"context"
	"github.com/explodes/serving/statusz"
)

var _ LogzServiceServer = (*logzServer)(nil)

var (
	varRecord = statusz.NewRateTracker("Record")
)

func registerLogzStatusz() {
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
	registerLogzStatusz()
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

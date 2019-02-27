package logz

import "context"

var _ LogzServiceServer = (*logzServer)(nil)

type logzServer struct {
	backend Backend
}

func NewLogzServer(backend Backend) LogzServiceServer {
	return &logzServer{
		backend: backend,
	}
}

func (s *logzServer) Record(ctx context.Context, req *RecordRequest) (*RecordResponse, error) {
	if err := s.backend.Record(req); err != nil {
		return nil, err
	}
	res := &RecordResponse{}
	return res, nil
}

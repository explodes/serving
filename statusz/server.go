package statusz

import "golang.org/x/net/context"

type statuszServer struct{}

func NewStatuszServer() StatuszServiceServer {
	return &statuszServer{}
}

func (s *statuszServer) GetStatusz(ctx context.Context, req *GetStatuszRequest) (*GetStatuszResponse, error) {
	res := &GetStatuszResponse{
		Statusz:&Statusz{},
	}
	return res, nil
}

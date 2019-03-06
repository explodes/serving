package utilz

import (
	"github.com/explodes/serving/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// DialGrpc dials a client connection taking into account dialing options.
func DialGrpc(addr *proto.GrpcServer) (*grpc.ClientConn, error) {
	if addr == nil || addr.Address == nil {
		return nil, errors.New("address not specified")
	}
	conn, err := grpc.Dial(addr.Address.Address(), grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "grpc dial error")
	}
	return conn, err
}

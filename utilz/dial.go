package utilz

import (
	spb "github.com/explodes/serving/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// DialGrpc dials a client connection taking into account dialing options.
func DialGrpc(serverConfig *spb.GrpcServer) (*grpc.ClientConn, error) {
	if serverConfig == nil || serverConfig.Address == nil {
		return nil, errors.New("address not specified")
	}
	opts := grpcDialOptions(serverConfig)
	conn, err := grpc.Dial(serverConfig.Address.Address(), opts...)
	if err != nil {
		return nil, errors.Wrap(err, "grpc dial error")
	}
	return conn, err
}

// grpcDialOptions extracts dial options from the GrpcServer config.
func grpcDialOptions(serverConfig *spb.GrpcServer) []grpc.DialOption {
	// TODO(evanleis): Add more features to GrpcServer.
	// Desired features:
	//  - backoff strategies
	//  - security
	//  - compression
	return []grpc.DialOption{grpc.WithInsecure()}
}

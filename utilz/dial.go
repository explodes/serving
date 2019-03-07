package utilz

import (
	"compress/gzip"
	"fmt"
	spb "github.com/explodes/serving/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	pbgzip "google.golang.org/grpc/encoding/gzip"
	"time"
)

const (
	defaultMaxDelay = 20 * time.Second

	defaultGzipCompressionLevel = gzip.BestSpeed
)

func init() {
	if err := pbgzip.SetLevel(defaultGzipCompressionLevel); err != nil {
		panic(fmt.Errorf("unable to set gzip compression level: %v", err))
	}
}

// DialGrpc dials a client connection taking into account dialing options.
func DialGrpc(config *spb.GrpcServer) (*grpc.ClientConn, error) {
	if config == nil || config.Address == nil {
		return nil, errors.New("address not specified")
	}
	opts := grpcDialOptions(config)
	conn, err := grpc.Dial(config.Address.Address(), opts...)
	if err != nil {
		return nil, errors.Wrap(err, "grpc dial error")
	}
	return conn, err
}

// grpcDialOptions extracts dial options from the GrpcServer config.
func grpcDialOptions(config *spb.GrpcServer) []grpc.DialOption {
	opts := make([]grpc.DialOption, 0, 3)

	// TODO(evanleis): security config
	opts = append(opts, grpc.WithInsecure())

	// Backoff strategy
	if backoff := config.ExponentialBackoff; backoff != nil {
		var maxDelay time.Duration
		if overrideDelay := backoff.OverrideMaxDelay; overrideDelay != nil {
			maxDelay = overrideDelay.Duration()
		} else {
			maxDelay = defaultMaxDelay
		}
		backoffConfig := grpc.BackoffConfig{MaxDelay: maxDelay}
		opts = append(opts, grpc.WithBackoffConfig(backoffConfig))
	}

	// Default call options, such as gzip.
	callOpts := grpcDefaultCallOptions(config)
	if len(callOpts) > 0 {
		opts = append(opts, grpc.WithDefaultCallOptions(callOpts...))
	}

	return opts
}

// grpcDefaultCallOptions creates default grpc.CallOptions from a config.
func grpcDefaultCallOptions(config *spb.GrpcServer) []grpc.CallOption {
	if config.EnableGzip {
		return []grpc.CallOption{grpc.UseCompressor(pbgzip.Name)}
	}
	return nil
}

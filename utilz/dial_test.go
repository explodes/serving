package utilz

import (
	spb "github.com/explodes/serving/proto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func sampleConfig() *spb.GrpcServer {
	return &spb.GrpcServer{
		EnableGzip: true,
		Address: &spb.Address{
			Host: "localhost",
			Port: 999,
		},
		ExponentialBackoff: &spb.GrpcServer_ExponentialBackoff{
			OverrideMaxDelay: spb.DurationSeconds(20),
		},
	}
}

func TestGrpcDialOptions(t *testing.T) {
	config := sampleConfig()

	opts := grpcDialOptions(config)

	assert.Len(t, opts, 3)
}

func TestGrpcDialOptions_noGzip(t *testing.T) {
	config := sampleConfig()
	config.EnableGzip = false

	opts := grpcDialOptions(config)

	assert.Len(t, opts, 2)
}

func TestGrpcDialOptions_noOverrideBackoff(t *testing.T) {
	config := sampleConfig()
	config.ExponentialBackoff.OverrideMaxDelay = nil

	opts := grpcDialOptions(config)

	assert.Len(t, opts, 3)
}

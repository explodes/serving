package logz_test

import (
	"github.com/explodes/serving/logz"
	"github.com/explodes/serving/test_serving"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newFrame() *logz.Frame {
	return &logz.Frame{
		FrameId:       "bar",
		ParentFrameId: "baz",
		FrameName:     "bell",
		StackId:       "bang",
	}
}

func TestServer(t *testing.T) {
	serverTest(t, "sanity", func(t *testing.T, server logz.LogzServiceServer, req *logz.RecordRequest) {
		res, err := server.Record(test_serving.TestContext(), req)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})
}

func serverTest(t *testing.T, name string, f func(t *testing.T, server logz.LogzServiceServer, req *logz.RecordRequest)) {
	t.Run(name, func(t *testing.T) {
		server := logz.NewLogzServer(nil)
		req := &logz.RecordRequest{Cookie: "some cookie", Frame: newFrame()}
		f(t, server, req)
	})
}

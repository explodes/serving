package expz_test

import (
	"github.com/explodes/serving/expz"
	"github.com/explodes/serving/expz/test_expz"
	"github.com/explodes/serving/logz/mock_logz"
	"github.com/explodes/serving/logz/test_logz"
	"github.com/explodes/serving/test_serving"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServer(t *testing.T) {
	serverTest(t, "sanity", func(t *testing.T, mockLogz *mock_logz.MockClient, server expz.ExpzServiceServer, req *expz.GetExperimentsRequest) {
		res, err := server.GetExperiments(test_serving.TestContext(), req)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})
}

func serverTest(t *testing.T, name string, f func(t *testing.T, mockLogz *mock_logz.MockClient, server expz.ExpzServiceServer, req *expz.GetExperimentsRequest)) {
	t.Run(name, func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockLogz := test_logz.NoopLogzClient(ctrl)
		server := expz.NewExpzServer(mockLogz, test_expz.ValidModFlags(t))
		req := &expz.GetExperimentsRequest{Cookie: "somecookie"}
		f(t, mockLogz, server, req)
	})
}

package addz_test

import (
	"errors"
	"github.com/explodes/serving/addz"
	"github.com/explodes/serving/expz/mock_expz"
	"github.com/explodes/serving/logz/test_logz"
	"github.com/explodes/serving/test_serving"
	"github.com/explodes/serving/userz/mock_userz"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddzServer_Add(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := test_serving.TestContext()
	expc := mock_expz.NewSettableMockClient()
	logc := test_logz.NoopLogzClient(ctrl)
	userc := mock_userz.NewMockClient(ctrl)
	userc.EXPECT().Validate(gomock.Any(), "somecookie").Return(true, nil)
	server := addz.NewAddzServer(logc, expc, userc)
	req := &addz.AddRequest{Cookie: "somecookie", Values: []int64{1, 2}}

	res, err := server.Add(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, int64(3), res.Result)
}

func TestAddzServer_Add_expz_error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := test_serving.TestContext()
	expc := mock_expz.NewSettableMockClient()
	expc.SetError(errors.New("mock expz error"))
	logc := test_logz.NoopLogzClient(ctrl)
	userc := mock_userz.NewMockClient(ctrl)
	userc.EXPECT().Validate(gomock.Any(), "somecookie").Return(true, nil)
	server := addz.NewAddzServer(logc, expc, userc)
	req := &addz.AddRequest{Cookie: "somecookie", Values: []int64{1, 2}}

	res, err := server.Add(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestAddzServer_Add_validate_error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := test_serving.TestContext()
	expc := mock_expz.NewSettableMockClient()
	logc := test_logz.NoopLogzClient(ctrl)
	userc := mock_userz.NewMockClient(ctrl)
	userc.EXPECT().Validate(gomock.Any(), "somecookie").Return(false, errors.New("validate error"))
	server := addz.NewAddzServer(logc, expc, userc)
	req := &addz.AddRequest{Cookie: "somecookie", Values: []int64{1, 2}}

	res, err := server.Add(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestAddzServer_Add_extra_addition(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := test_serving.TestContext()
	expc := mock_expz.NewSettableMockClient()
	expc.SetFlagI64("extra_addition", 100)
	logc := test_logz.NoopLogzClient(ctrl)
	userc := mock_userz.NewMockClient(ctrl)
	userc.EXPECT().Validate(gomock.Any(), "somecookie").Return(true, nil)
	server := addz.NewAddzServer(logc, expc, userc)
	req := &addz.AddRequest{Cookie: "somecookie", Values: []int64{1, 2}}

	res, err := server.Add(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, int64(103), res.Result)
}

func TestAddzServer_Add_Anon(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := test_serving.TestContext()
	expc := mock_expz.NewSettableMockClient()
	logc := test_logz.NoopLogzClient(ctrl)
	userc := mock_userz.NewMockClient(ctrl)
	userc.EXPECT().Validate(gomock.Any(), "somecookie").Return(false, nil)
	server := addz.NewAddzServer(logc, expc, userc)
	req := &addz.AddRequest{Cookie: "somecookie", Values: []int64{1, 2}}

	res, err := server.Add(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestAddzServer_Subtract(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := test_serving.TestContext()
	expc := mock_expz.NewSettableMockClient()
	logc := test_logz.NoopLogzClient(ctrl)
	userc := mock_userz.NewMockClient(ctrl)
	userc.EXPECT().Validate(gomock.Any(), "somecookie").Return(true, nil)
	server := addz.NewAddzServer(logc, expc, userc)
	req := &addz.SubtractRequest{Cookie: "somecookie", Values: []int64{1, 2}}

	res, err := server.Subtract(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, int64(-1), res.Result)
}

func TestAddzServer_Subtract_noValues(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := test_serving.TestContext()
	expc := mock_expz.NewSettableMockClient()
	logc := test_logz.NoopLogzClient(ctrl)
	userc := mock_userz.NewMockClient(ctrl)
	userc.EXPECT().Validate(gomock.Any(), "somecookie").Return(true, nil)
	server := addz.NewAddzServer(logc, expc, userc)
	req := &addz.SubtractRequest{Cookie: "somecookie"}

	res, err := server.Subtract(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, int64(0), res.Result)
}

func TestAddzServer_Subtract_Anon(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := test_serving.TestContext()
	expc := mock_expz.NewSettableMockClient()
	logc := test_logz.NoopLogzClient(ctrl)
	userc := mock_userz.NewMockClient(ctrl)
	userc.EXPECT().Validate(gomock.Any(), "somecookie").Return(false, nil)
	server := addz.NewAddzServer(logc, expc, userc)
	req := &addz.SubtractRequest{Cookie: "somecookie", Values: []int64{1, 2}}

	res, err := server.Subtract(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestAddzServer_Subtract_expz_error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := test_serving.TestContext()
	expc := mock_expz.NewSettableMockClient()
	expc.SetError(errors.New("mock expz error"))
	logc := test_logz.NoopLogzClient(ctrl)
	userc := mock_userz.NewMockClient(ctrl)
	userc.EXPECT().Validate(gomock.Any(), "somecookie").Return(true, nil)
	server := addz.NewAddzServer(logc, expc, userc)
	req := &addz.SubtractRequest{Cookie: "somecookie", Values: []int64{1, 2}}

	res, err := server.Subtract(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestAddzServer_Subtract_validate_error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := test_serving.TestContext()
	expc := mock_expz.NewSettableMockClient()
	logc := test_logz.NoopLogzClient(ctrl)
	userc := mock_userz.NewMockClient(ctrl)
	userc.EXPECT().Validate(gomock.Any(), "somecookie").Return(false, errors.New("validate error"))
	server := addz.NewAddzServer(logc, expc, userc)
	req := &addz.SubtractRequest{Cookie: "somecookie", Values: []int64{1, 2}}

	res, err := server.Subtract(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, res)
}

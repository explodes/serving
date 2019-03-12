package userz_test

import (
	"errors"
	"github.com/explodes/serving/expz/mock_expz"
	"github.com/explodes/serving/logz/test_logz"
	"github.com/explodes/serving/test_serving"
	"github.com/explodes/serving/userz"
	"github.com/explodes/serving/userz/mock_userz"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUserzServer_Login_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := test_serving.TestContext()
	expz := mock_expz.NewSettableMockClient()
	logz := test_logz.NoopLogzClient(ctrl)
	storage := mock_userz.NewMockStorage(ctrl)
	storage.EXPECT().UserID(ctx, "username", "password").Return("userid", nil)
	storage.EXPECT().Save(ctx, gomock.Any(), gomock.Any()).Return(nil)
	server := userz.NewUserzServer("cookiepass", storage, logz, expz)
	req := &userz.LoginRequest{Username: "username", Password: "password"}

	res, err := server.Login(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.True(t, len(res.Result.(*userz.LoginResponse_Success).Success.Cookie) > 32)
}

func TestNewUserzServer_Login_BadLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := test_serving.TestContext()
	expz := mock_expz.NewSettableMockClient()
	logz := test_logz.NoopLogzClient(ctrl)
	storage := mock_userz.NewMockStorage(ctrl)
	storage.EXPECT().UserID(ctx, "username", "password").Return("", errors.New("invalid"))
	server := userz.NewUserzServer("cookiepass", storage, logz, expz)
	req := &userz.LoginRequest{Username: "username", Password: "password"}

	res, err := server.Login(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, userz.LoginResponse_LoginFailure_BAD_LOGIN, res.Result.(*userz.LoginResponse_Failure).Failure.Reason)
}

func TestNewUserzServer_Login_SaveFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := test_serving.TestContext()
	expz := mock_expz.NewSettableMockClient()
	logz := test_logz.NoopLogzClient(ctrl)
	storage := mock_userz.NewMockStorage(ctrl)
	storage.EXPECT().UserID(ctx, "username", "password").Return("userid", nil)
	storage.EXPECT().Save(ctx, gomock.Any(), gomock.Any()).Return(errors.New("simulated save error"))
	server := userz.NewUserzServer("cookiepass", storage, logz, expz)
	req := &userz.LoginRequest{Username: "username", Password: "password"}

	res, err := server.Login(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestNewUserzServer_Validate_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := test_serving.TestContext()
	expz := mock_expz.NewSettableMockClient()
	logz := test_logz.NoopLogzClient(ctrl)
	storage := mock_userz.NewMockStorage(ctrl)
	storage.EXPECT().Validate(ctx, "cookie").Return(true, nil)
	server := userz.NewUserzServer("cookiepass", storage, logz, expz)
	req := &userz.ValidateRequest{Cookie: "cookie"}

	res, err := server.Validate(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, userz.ValidateResponse_SUCCESS, res.Result)
}

func TestNewUserzServer_Validate_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := test_serving.TestContext()
	expz := mock_expz.NewSettableMockClient()
	logz := test_logz.NoopLogzClient(ctrl)
	storage := mock_userz.NewMockStorage(ctrl)
	storage.EXPECT().Validate(ctx, "cookie").Return(false, nil)
	server := userz.NewUserzServer("cookiepass", storage, logz, expz)
	req := &userz.ValidateRequest{Cookie: "cookie"}

	res, err := server.Validate(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, userz.ValidateResponse_INVALID, res.Result)
}

func TestNewUserzServer_Validate_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := test_serving.TestContext()
	expz := mock_expz.NewSettableMockClient()
	logz := test_logz.NoopLogzClient(ctrl)
	storage := mock_userz.NewMockStorage(ctrl)
	storage.EXPECT().Validate(ctx, "cookie").Return(false, errors.New("simulated validate error"))
	server := userz.NewUserzServer("cookiepass", storage, logz, expz)
	req := &userz.ValidateRequest{Cookie: "cookie"}

	res, err := server.Validate(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, res)
}

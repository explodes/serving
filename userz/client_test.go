package userz_test

import (
	"errors"
	"github.com/explodes/serving/test_serving"
	"github.com/explodes/serving/userz"
	"github.com/explodes/serving/userz/mock_userz"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClientImpl_Validate_success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserzServiceClient := mock_userz.NewMockUserzServiceClient(ctrl)
	ctx := test_serving.TestContext()
	mockUserzServiceClient.EXPECT().Validate(ctx, gomock.Any()).Return(&userz.ValidateResponse{Result: userz.ValidateResponse_SUCCESS}, nil)
	client := userz.NewClient(mockUserzServiceClient)

	valid, err := client.Validate(ctx, "cookie")

	assert.NoError(t, err)
	assert.True(t, valid)
}

func TestClientImpl_Validate_invalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserzServiceClient := mock_userz.NewMockUserzServiceClient(ctrl)
	ctx := test_serving.TestContext()
	mockUserzServiceClient.EXPECT().Validate(ctx, gomock.Any()).Return(&userz.ValidateResponse{Result: userz.ValidateResponse_INVALID}, nil)
	client := userz.NewClient(mockUserzServiceClient)

	valid, err := client.Validate(ctx, "cookie")

	assert.NoError(t, err)
	assert.False(t, valid)
}

func TestClientImpl_Validate_error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserzServiceClient := mock_userz.NewMockUserzServiceClient(ctrl)
	ctx := test_serving.TestContext()
	mockUserzServiceClient.EXPECT().Validate(ctx, gomock.Any()).Return(nil, errors.New("some error"))
	client := userz.NewClient(mockUserzServiceClient)

	_, err := client.Validate(ctx, "cookie")

	assert.Error(t, err)
}

func TestClientImpl_Login_ok(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserzServiceClient := mock_userz.NewMockUserzServiceClient(ctrl)
	ctx := test_serving.TestContext()
	mockUserzServiceClient.EXPECT().Login(ctx, gomock.Any()).Return(&userz.LoginResponse{Result: &userz.LoginResponse_Success{Success: &userz.LoginResponse_LoginSuccess{Cookie: "cookie"}}}, nil)
	client := userz.NewClient(mockUserzServiceClient)

	cookie, err := client.Login(ctx, "test", "pass")

	assert.NoError(t, err)
	assert.Equal(t, "cookie", cookie)
}

func TestClientImpl_Login_invalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserzServiceClient := mock_userz.NewMockUserzServiceClient(ctrl)
	ctx := test_serving.TestContext()
	mockUserzServiceClient.EXPECT().Login(ctx, gomock.Any()).Return(&userz.LoginResponse{Result: &userz.LoginResponse_Failure{Failure: &userz.LoginResponse_LoginFailure{Reason: userz.LoginResponse_LoginFailure_BAD_LOGIN}}}, nil)
	client := userz.NewClient(mockUserzServiceClient)

	cookie, err := client.Login(ctx, "test", "pass")

	assert.Error(t, err)
	assert.Equal(t, "", cookie)
}

func TestClientImpl_Login_error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserzServiceClient := mock_userz.NewMockUserzServiceClient(ctrl)
	ctx := test_serving.TestContext()
	mockUserzServiceClient.EXPECT().Login(ctx, gomock.Any()).Return(nil, errors.New("some error"))
	client := userz.NewClient(mockUserzServiceClient)

	cookie, err := client.Login(ctx, "test", "pass")

	assert.Error(t, err)
	assert.Equal(t, "", cookie)
}

func TestClientImpl_Login_unknownFailureType(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserzServiceClient := mock_userz.NewMockUserzServiceClient(ctrl)
	ctx := test_serving.TestContext()
	mockUserzServiceClient.EXPECT().Login(ctx, gomock.Any()).Return(&userz.LoginResponse{Result: &userz.LoginResponse_Failure{Failure: &userz.LoginResponse_LoginFailure{Reason: 999}}}, nil)
	client := userz.NewClient(mockUserzServiceClient)

	cookie, err := client.Login(ctx, "test", "pass")

	assert.Error(t, err)
	assert.Equal(t, "", cookie)
}

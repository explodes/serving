package userz

import (
	"context"
	"github.com/pkg/errors"
)

type Client interface {
	Login(ctx context.Context, username, password string) (cookie string, err error)
	Validate(ctx context.Context, cookie string) (bool, error)
}

type clientImpl struct {
	userz UserzServiceClient
}

func NewClient(userz UserzServiceClient) Client {
	client := &clientImpl{
		userz: userz,
	}
	return client
}

func (c *clientImpl) Login(ctx context.Context, username, password string) (cookie string, err error) {
	req := &LoginRequest{
		Username: username,
		Password: password,
	}
	res, err := c.userz.Login(ctx, req)
	if err != nil {
		return "", errors.Wrap(err, "could not login")
	}
	switch t := res.Result.(type) {
	case *LoginResponse_Success:
		return t.Success.Cookie, nil
	case *LoginResponse_Failure:
		switch t.Failure.Reason {
		case LoginResponse_LoginFailure_BAD_LOGIN:
			return "", errors.New("invalid login")
		default:
			return "", errors.New("unknown login failure")
		}
	default:
		return "", errors.New("unknown login result")
	}
}

func (c *clientImpl) Validate(ctx context.Context, cookie string) (bool, error) {
	res, err := c.userz.Validate(ctx, &ValidateRequest{Cookie: cookie})
	if err != nil {
		return false, errors.Wrap(err, "error validating user")
	}
	return res.Result == ValidateResponse_SUCCESS, nil
}

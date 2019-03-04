package userz

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"sync"
)

type Client struct {
	clientMu *sync.RWMutex
	addr     string
	conn     *grpc.ClientConn
	userz    UserzServiceClient
}

func NewClient(addr string) (*Client, error) {
	client := &Client{
		clientMu: &sync.RWMutex{},
		addr:     addr,
	}
	err := client.restoreClient()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) restoreClient() error {
	c.clientMu.Lock()
	defer c.clientMu.Unlock()
	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			return err
		}
		c.conn = nil
		c.userz = nil
	}
	conn, err := grpc.Dial(c.addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	c.conn = conn
	c.userz = NewUserzServiceClient(conn)
	return nil
}

func (c *Client) Login(ctx context.Context, username, password string) (cookie string, err error) {
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

func (c *Client) Validate(ctx context.Context, cookie string) (bool, error) {
	res, err := c.userz.Validate(ctx, &ValidateRequest{Cookie: cookie})
	if err != nil {
		return false, errors.Wrap(err, "error validating user")
	}
	return res.Result == ValidateResponse_SUCCESS, nil
}

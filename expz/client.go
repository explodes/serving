package expz

import (
	"context"
	spb "github.com/explodes/serving/proto"
	"google.golang.org/grpc"
	"sync"
)

type Client struct {
	clientMu *sync.RWMutex
	addr     *spb.Address
	conn     *grpc.ClientConn
	expz     ExpzServiceClient
}

func NewClient(addr *spb.Address) (*Client, error) {
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
		c.expz = nil
	}
	conn, err := grpc.Dial(c.addr.Address(), grpc.WithInsecure())
	if err != nil {
		return err
	}
	c.conn = conn
	c.expz = NewExpzServiceClient(conn)
	return nil
}

func (c *Client) GetExperiments(ctx context.Context, cookie int64) (*ExperimentFlags, error) {
	req := &GetExperimentsRequest{Cookie: cookie}
	res, err := c.expz.GetExperiments(ctx, req)
	return &ExperimentFlags{res: res}, err
}

type ExperimentFlags struct {
	res *GetExperimentsResponse
}

func (exp ExperimentFlags) GetFlag(name string) *Flag {
	flag, ok := exp.res.Features.Flags[name]
	if !ok {
		return nil
	}
	return flag
}

func (exp ExperimentFlags) Int64Value(name string, def int64) int64 {
	flag := exp.GetFlag(name)
	if flag == nil {
		return def
	}
	return flag.Int64Value(def)
}

func (exp ExperimentFlags) Float64Value(name string, def float64) float64 {
	flag := exp.GetFlag(name)
	if flag == nil {
		return def
	}
	return flag.Float64Value(def)
}

func (exp ExperimentFlags) BoolValue(name string, def bool) bool {
	flag := exp.GetFlag(name)
	if flag == nil {
		return def
	}
	return flag.BoolValue(def)
}

func (exp ExperimentFlags) StringValue(name string, def string) string {
	flag := exp.GetFlag(name)
	if flag == nil {
		return def
	}
	return flag.StringValue(def)
}

package expz

import (
	"context"
)

type Client struct {
	expz ExpzServiceClient
}

func NewClient(expz ExpzServiceClient) *Client {
	client := &Client{
		expz: expz,
	}
	return client
}

func (c *Client) GetExperiments(ctx context.Context, cookie string) (*ExperimentFlags, error) {
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

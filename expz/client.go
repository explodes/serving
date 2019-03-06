package expz

import (
	"context"
)

type Client interface {
	GetExperiments(ctx context.Context, cookie string) (ExperimentFlags, error)
}

type clientImpl struct {
	expz ExpzServiceClient
}

func NewClient(expz ExpzServiceClient) Client {
	client := &clientImpl{
		expz: expz,
	}
	return client
}

func (c *clientImpl) GetExperiments(ctx context.Context, cookie string) (ExperimentFlags, error) {
	req := &GetExperimentsRequest{Cookie: cookie}
	res, err := c.expz.GetExperiments(ctx, req)
	return NewExperimentFlags(res.Features.Flags), err
}

type ExperimentFlags map[string]*Flag

func NewExperimentFlags(m map[string]*Flag) ExperimentFlags {
	return ExperimentFlags(m)
}

func (ef ExperimentFlags) GetFlag(name string) *Flag {
	flag, ok := ef[name]
	if !ok {
		return nil
	}
	return flag
}

func (ef ExperimentFlags) Int64Value(name string, def int64) int64 {
	flag := ef.GetFlag(name)
	if flag == nil {
		return def
	}
	return flag.Int64Value(def)
}

func (ef ExperimentFlags) Float64Value(name string, def float64) float64 {
	flag := ef.GetFlag(name)
	if flag == nil {
		return def
	}
	return flag.Float64Value(def)
}

func (ef ExperimentFlags) BoolValue(name string, def bool) bool {
	flag := ef.GetFlag(name)
	if flag == nil {
		return def
	}
	return flag.BoolValue(def)
}

func (ef ExperimentFlags) StringValue(name string, def string) string {
	flag := ef.GetFlag(name)
	if flag == nil {
		return def
	}
	return flag.StringValue(def)
}

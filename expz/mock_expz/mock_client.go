package mock_expz

import (
	"context"
	"github.com/explodes/serving/expz"
)

var _ expz.Client = (*SettableMockClient)(nil)

type SettableMockClient struct {
	m map[string]*expz.Flag
}

func NewSettableMockClient() *SettableMockClient {
	return &SettableMockClient{
		m: make(map[string]*expz.Flag),
	}
}

func (m *SettableMockClient) SetFlag(name string, flag *expz.Flag) *SettableMockClient {
	m.m[name] = flag
	return m
}

func (m *SettableMockClient) SetFlagI64(name string, value int64) *SettableMockClient {
	return m.SetFlag(name, &expz.Flag{Flag: &expz.Flag_I64{I64: value}})
}

func (m *SettableMockClient) SetFlagU64(name string, value float64) *SettableMockClient {
	return m.SetFlag(name, &expz.Flag{Flag: &expz.Flag_F64{F64: value}})
}

func (m *SettableMockClient) SetFlagString(name string, value string) *SettableMockClient {
	return m.SetFlag(name, &expz.Flag{Flag: &expz.Flag_String_{String_: value}})
}

func (m *SettableMockClient) SetFlagBool(name string, value bool) *SettableMockClient {
	return m.SetFlag(name, &expz.Flag{Flag: &expz.Flag_Bool{Bool: value}})
}

func (m *SettableMockClient) GetExperiments(ctx context.Context, cookie string) (expz.ExperimentFlags, error) {
	return m.m, nil
}

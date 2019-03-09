package statusz_test

import (
	"errors"
	"github.com/explodes/serving/statusz"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVarGroup_MarshalMetrics(t *testing.T) {
	vg := statusz.VarGroup{
		noopVar{},
		noopVarMetric{},
	}
	metrics, err := vg.MarshalMetrics()
	assert.NoError(t, err)
	assert.NotNil(t, metrics)
	assert.Len(t, metrics, 2)
}

func TestVarGroup_MarshalMetrics_errVar(t *testing.T) {
	vg := statusz.VarGroup{
		errVar{err: errors.New("test")},
	}
	metrics, err := vg.MarshalMetrics()
	assert.Error(t, err)
	assert.Nil(t, metrics)
}

func TestVarGroup_MarshalMetrics_errVarMetric(t *testing.T) {
	vg := statusz.VarGroup{
		errVarMetric{err: errors.New("test")},
	}
	metrics, err := vg.MarshalMetrics()
	assert.Error(t, err)
	assert.Nil(t, metrics)
}

func TestVarGroup_MarshalMetrics_UnknownType(t *testing.T) {
	vg := statusz.VarGroup{
		1,
	}
	defer func() {
		assert.NotNil(t, recover())
	}()
	metrics, err := vg.MarshalMetrics()
	assert.Fail(t, "should have panicked: %v, %v", metrics, err)

}

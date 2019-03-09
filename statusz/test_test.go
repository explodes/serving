package statusz_test

import (
	"errors"
	"github.com/explodes/serving/statusz"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var _ statusz.Var = noopVar{}
var _ statusz.Var = errVar{}
var _ statusz.VarMetric = noopVarMetric{}
var _ statusz.VarMetric = errVarMetric{}

type noopVar struct{}

func (noopVar) MarshalMetrics() ([]*statusz.Metric, error) {
	return []*statusz.Metric{someMetric()}, nil
}

type errVar struct {
	err error
}

func (e errVar) MarshalMetrics() ([]*statusz.Metric, error) {
	return nil, error(e.err)
}

type noopVarMetric struct{}

func (noopVarMetric) MarshalMetric() (*statusz.Metric, error) {
	return someMetric(), nil
}

type errVarMetric struct {
	err error
}

func (e errVarMetric) MarshalMetric() (*statusz.Metric, error) {
	return nil, error(e.err)
}

func someMetric() *statusz.Metric {
	return &statusz.Metric{
		Name: "foo",
		Value: &statusz.Metric_I64{
			I64: 100,
		},
	}
}

type brokenPipe struct {
	statusCode int
}

func (b *brokenPipe) Header() http.Header {
	return make(http.Header)
}

func (b *brokenPipe) Write([]byte) (int, error) {
	return 0, errors.New("broken pipe")
}

func (b *brokenPipe) WriteHeader(statusCode int) {
	b.statusCode = statusCode
}

func readBody(t *testing.T, w *httptest.ResponseRecorder) string {
	t.Helper()

	b, err := ioutil.ReadAll(w.Result().Body)
	assert.NoError(t, err)

	return string(b)
}
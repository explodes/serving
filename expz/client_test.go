package expz_test

import (
	"context"
	"github.com/explodes/serving/expz"
	"github.com/explodes/serving/expz/mock_expz"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func mockGetExperimentsResponse() *expz.GetExperimentsResponse {
	return &expz.GetExperimentsResponse{
		Features: &expz.Features{
			Flags: map[string]*expz.Flag{
				"string_flag": {Flag: &expz.Flag_String_{String_: "value"}},
				"i64_flag":    {Flag: &expz.Flag_I64{I64: 100}},
				"f64_flag":    {Flag: &expz.Flag_F64{F64: 200.}},
				"bool_flag":   {Flag: &expz.Flag_Bool{Bool: true}},
			},
		},
	}
}

func TestClient_GetExperiments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := mock_expz.NewMockExpzServiceClient(ctrl)
	mock.EXPECT().GetExperiments(gomock.Any(), gomock.Any()).Return(mockGetExperimentsResponse(), nil)
	client := expz.NewClient(mock)

	exps, err := client.GetExperiments(context.Background(), "test")
	assert.NoError(t, err)

	// Expected flag values.
	assert.Equal(t, "value", exps.StringValue("string_flag", "unexpected"))
	assert.Equal(t, int64(100), exps.Int64Value("i64_flag", 783945))
	assert.Equal(t, float64(200), exps.Float64Value("f64_flag", 900.))
	assert.Equal(t, true, exps.BoolValue("bool_flag", false))

	// Default values for flags that do not exist.
	assert.Equal(t, "expected", exps.StringValue("nonexistent_string_flag", "expected"))
	assert.Equal(t, int64(783945), exps.Int64Value("nonexistent_i64_flag", 783945))
	assert.Equal(t, float64(900), exps.Float64Value("nonexistent_f64_flag", 900.))
	assert.Equal(t, false, exps.BoolValue("nonexistent_bool_flag", false))
}

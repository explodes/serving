// +build testing

package test_expz

import (
	"github.com/explodes/serving/expz"
	"github.com/stretchr/testify/assert"
	"testing"
)

func ValidConfig() *expz.ExpzConfig {
	return &expz.ExpzConfig{
		DefaultFeatures: []*expz.FeatureDeclaration{
			{
				Name: "myflag",
				Doc:  "mydoc",
				DefaultValue: &expz.Flag{
					Flag: &expz.Flag_Bool{
						Bool: false,
					},
				},
			},
		},
		ExperimentalFeatures: []*expz.ExpzConfig_ExperimentalFeatures{
			{
				Name: "myexperiment",
				Doc:  "someexperiment",
				Mods: &expz.ExpzConfig_ExperimentalFeatures_Mods{
					Min: 0,
					Max: 999,
				},
				Features: &expz.Features{
					Flags: map[string]*expz.Flag{
						"myflag": &expz.Flag{
							Flag: &expz.Flag_Bool{
								Bool: true,
							},
						},
					},
				},
			},
		},
	}
}

func ValidModFlags(t *testing.T) expz.ModFlags {
	t.Helper()
	mods, err := ValidConfig().Validate()
	assert.NoError(t, err)
	return mods
}

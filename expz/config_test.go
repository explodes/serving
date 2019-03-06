package expz_test

import (
	"github.com/explodes/serving/expz"
	"github.com/explodes/serving/expz/test_expz"
	"github.com/stretchr/testify/assert"
	"testing"
)

func assertConfig(t *testing.T, hasErr bool, name string, mutator func(config *expz.ExpzConfig, feature *expz.FeatureDeclaration, experiment *expz.ExpzConfig_ExperimentalFeatures)) {
	t.Run(name, func(t *testing.T) {
		config := test_expz.ValidConfig()
		feature := config.DefaultFeatures[0]
		experiment := config.ExperimentalFeatures[0]
		mutator(config, feature, experiment)
		_, err := config.Validate()
		if hasErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	})
}

func assertInvalid(t *testing.T, name string, mutator func(config *expz.ExpzConfig, feature *expz.FeatureDeclaration, experiment *expz.ExpzConfig_ExperimentalFeatures)) {
	assertConfig(t, true, name, mutator)
}

func assertValid(t *testing.T, name string, mutator func(config *expz.ExpzConfig, feature *expz.FeatureDeclaration, experiment *expz.ExpzConfig_ExperimentalFeatures)) {
	assertConfig(t, false, name, mutator)
}

func TestExpzConfig_Validate(t *testing.T) {
	assertValid(t, "valid", func(config *expz.ExpzConfig, feature *expz.FeatureDeclaration, experiment *expz.ExpzConfig_ExperimentalFeatures) {})
	assertValid(t, "empty", func(config *expz.ExpzConfig, feature *expz.FeatureDeclaration, experiment *expz.ExpzConfig_ExperimentalFeatures) {
		config.DefaultFeatures = nil
		config.ExperimentalFeatures = nil
	})
	assertValid(t, "no_experiments", func(config *expz.ExpzConfig, feature *expz.FeatureDeclaration, experiment *expz.ExpzConfig_ExperimentalFeatures) {
		config.ExperimentalFeatures = nil
	})
	assertValid(t, "experiment_without_features", func(config *expz.ExpzConfig, feature *expz.FeatureDeclaration, experiment *expz.ExpzConfig_ExperimentalFeatures) {
		experiment.Features = nil
	})
	assertInvalid(t, "feature_no_name", func(config *expz.ExpzConfig, feature *expz.FeatureDeclaration, experiment *expz.ExpzConfig_ExperimentalFeatures) {
		feature.Name = ""
	})
	assertInvalid(t, "feature_no_doc", func(config *expz.ExpzConfig, feature *expz.FeatureDeclaration, experiment *expz.ExpzConfig_ExperimentalFeatures) {
		feature.Doc = ""
	})
	assertInvalid(t, "feature_duplicate", func(config *expz.ExpzConfig, feature *expz.FeatureDeclaration, experiment *expz.ExpzConfig_ExperimentalFeatures) {
		config.DefaultFeatures = append(config.DefaultFeatures, feature)
	})
	assertInvalid(t, "feature_no_default", func(config *expz.ExpzConfig, feature *expz.FeatureDeclaration, experiment *expz.ExpzConfig_ExperimentalFeatures) {
		feature.DefaultValue = nil
	})
	assertInvalid(t, "feature_no_default_flag", func(config *expz.ExpzConfig, feature *expz.FeatureDeclaration, experiment *expz.ExpzConfig_ExperimentalFeatures) {
		feature.DefaultValue.Flag = nil
	})
	assertInvalid(t, "experiment_no_name", func(config *expz.ExpzConfig, feature *expz.FeatureDeclaration, experiment *expz.ExpzConfig_ExperimentalFeatures) {
		experiment.Name = ""
	})
	assertInvalid(t, "experiment_no_doc", func(config *expz.ExpzConfig, feature *expz.FeatureDeclaration, experiment *expz.ExpzConfig_ExperimentalFeatures) {
		experiment.Doc = ""
	})
	assertInvalid(t, "experiment_no_mods", func(config *expz.ExpzConfig, feature *expz.FeatureDeclaration, experiment *expz.ExpzConfig_ExperimentalFeatures) {
		experiment.Mods = nil
	})
	assertInvalid(t, "experiment_mods_max", func(config *expz.ExpzConfig, feature *expz.FeatureDeclaration, experiment *expz.ExpzConfig_ExperimentalFeatures) {
		experiment.Mods.Max = expz.MaxMods
	})
	assertInvalid(t, "experiment_mods_min_gt_max", func(config *expz.ExpzConfig, feature *expz.FeatureDeclaration, experiment *expz.ExpzConfig_ExperimentalFeatures) {
		experiment.Mods.Max = 10
		experiment.Mods.Min = 11
	})
	assertInvalid(t, "experiment_flag_doesnt_exist", func(config *expz.ExpzConfig, feature *expz.FeatureDeclaration, experiment *expz.ExpzConfig_ExperimentalFeatures) {
		experiment.Features.Flags["nonexistent"] = &expz.Flag{
			Flag: &expz.Flag_Bool{
				Bool: true,
			},
		}
	})
	assertInvalid(t, "experiment_flag_wrong_type", func(config *expz.ExpzConfig, feature *expz.FeatureDeclaration, experiment *expz.ExpzConfig_ExperimentalFeatures) {
		experiment.Features.Flags["myflag"].Flag = &expz.Flag_I64{}
	})
	assertInvalid(t, "experiment_overrides", func(config *expz.ExpzConfig, feature *expz.FeatureDeclaration, experiment *expz.ExpzConfig_ExperimentalFeatures) {
		exp2 := &expz.ExpzConfig_ExperimentalFeatures{
			Name: "overriding",
			Doc:  "overriding",
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
		}
		config.ExperimentalFeatures = append(config.ExperimentalFeatures, exp2)
	})
}

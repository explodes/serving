package expz

import (
	"errors"
	"fmt"
)

var emptyExperiments = Experiments{}

func (m *ExpzConfig) Validate() (Experiments, error) {
	// Collect known flags.
	experiments := NewExperiments()
	definedFlags := make(map[string]flagType)
	experimentalFlags := NewExperiments()

	for _, flag := range m.DefaultFeatures {
		// Validate the FeatureDeclaration.
		if flag.Name == "" {
			return emptyExperiments, errors.New("default feature is missing a name")
		}
		if flag.Doc == "" {
			return emptyExperiments, fmt.Errorf("default feature %s is missing documentation", flag.Name)
		}
		if flag.DefaultValue == nil {
			return emptyExperiments, fmt.Errorf("default feature %s is missing a default value", flag.Name)
		}
		// Cannot be defined multiple times.
		if _, exists := definedFlags[flag.Name]; exists {
			return emptyExperiments, fmt.Errorf("%s flag defined multiple times", flag.Name)
		}
		// Record the type of the flag.
		fType, err := flag.DefaultValue.flagType()
		if err != nil {
			return emptyExperiments, fmt.Errorf("error determining type of flag %s: %v", flag.Name, err)
		}
		definedFlags[flag.Name] = fType
		// Store this value in our compressed experiments structure.
		experiments.setDefault(flag.Name, flag.DefaultValue)
	}
	// Validate experiment configuration against the set of known flags.
	for _, exp := range m.ExperimentalFeatures {
		if exp.Name == "" {
			return emptyExperiments, errors.New("an experiment is missing a name")
		}
		if exp.Doc == "" {
			return emptyExperiments, fmt.Errorf("experiment %s is missing documentation", exp.Name)
		}
		if exp.Mods == nil {
			return emptyExperiments, fmt.Errorf("experiment %s is missing mods", exp.Name)
		}
		if exp.Mods.Min < 0 || exp.Mods.Min > MaxMods || exp.Mods.Max < 0 || exp.Mods.Max >= MaxMods {
			return emptyExperiments, fmt.Errorf("experiment %s is mods out of range [0, %d]", exp.Name, MaxMods)
		}
		if exp.Mods.Max < exp.Mods.Min {
			return emptyExperiments, fmt.Errorf("experiment %s mods misconfigured: min must be less less than or equal to max", exp.Name)
		}
		for name, flag := range exp.Features.Flags {
			definedFlagType, exists := definedFlags[name]
			if !exists {
				return emptyExperiments, fmt.Errorf("experiment %s declares non-existant flag %s", exp.Name, name)
			}
			flagType, err := flag.flagType()
			if err != nil {
				return emptyExperiments, fmt.Errorf("error determining type of flag %s for experiment %s: %v", name, exp.Name, err)
			}
			if flagType != definedFlagType {
				return emptyExperiments, fmt.Errorf("experiment %s flag %s has type other than previously defined", exp.Name, name)
			}
			if experimentalFlags.rangeContains(name, int(exp.Mods.Min), int(exp.Mods.Max)) {
				return emptyExperiments, fmt.Errorf("experiment %s flag %s overrides flag defined in a mod overwritten by a previous experiment", exp.Name, name)
			}
			// Store this value in our compressed experiments structure.
			experiments.setRange(name, int(exp.Mods.Min), int(exp.Mods.Max), flag)
			// Store this value in our collection of experiment-overriden flags.
			experimentalFlags.setRange(name, int(exp.Mods.Min), int(exp.Mods.Max), flag)
		}
	}
	return experiments, nil
}

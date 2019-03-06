package expz

const (
	MaxMods = 1000
)

type ModFlags [MaxMods]map[string]*Flag

func NewModFlags() ModFlags {
	var exps ModFlags
	for index := range exps {
		exps[index] = make(map[string]*Flag)
	}
	return exps
}

// setDefault sets the default value for all mods.
func (e *ModFlags) setDefault(name string, defaultValue *Flag) {
	e.setRange(name, 0, MaxMods-1, defaultValue)
}

// setRange sets the default value of a flag for mods in an inclusive range.
func (e *ModFlags) setRange(name string, min, max int, defaultValue *Flag) {
	for index := min; index <= max; index++ {
		e[index][name] = defaultValue
	}
}

// rangeContains tests if mods in an inclusive range contains a definition
// for a flag with a given name.
func (e *ModFlags) rangeContains(name string, min, max int) bool {
	for index := min; index <= max; index++ {
		_, ok := e[index][name]
		if ok {
			return true
		}
	}
	return false
}

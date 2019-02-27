package statusz

import (
	"fmt"
	"sync"
)

var (
	varRegistryLock = &sync.Mutex{}
	varRegistry     = make(varRegistryMap, 0, 16)
)

func Register(name string, v Var) {
	varRegistryLock.Lock()
	if varRegistry.exists(name) {
		panic(fmt.Errorf("a var with name %s is already registered", name))
	}
	varRegistry = append(varRegistry, namedVar{name: name, v: v})
	varRegistryLock.Unlock()
}

type namedVar struct {
	name string
	v    Var
}

type varRegistryMap []namedVar

func (v varRegistryMap) exists(name string) bool {
	for _, named := range v {
		if named.name == name {
			return true
		}
	}
	return false
}

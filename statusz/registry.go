package statusz

import (
	"fmt"
	"sync"
)

var (
	varRegistryLock = &sync.Mutex{}
	varRegistry     = make(varRegistryMap, 0, 16)
)

type namedVar struct {
	name string
	v    Var
}

func Register(name string, v Var) {
	varRegistryLock.Lock()
	defer varRegistryLock.Unlock()
	if varRegistry.exists(name) {
		panic(fmt.Errorf("a var with name %s is already registered", name))
	}
	varRegistry = append(varRegistry, namedVar{name: name, v: v})

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

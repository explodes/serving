// +build testing

package statusz

func ResetVarz() {
	varRegistryLock.Lock()
	varRegistry = make(varRegistryMap, 0)
	varRegistryLock.Unlock()
}

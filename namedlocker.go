package namedlocker

import (
	"sync"
)

// NamedLocker implements locks by name with lazy object creation.
// Mutexes are implemented as RWMutex
type NamedLocker struct {
	locks *sync.Map
}

// NewNamedLocker creates a new named locker
func NewNamedLocker() *NamedLocker {
	return &NamedLocker{
		locks: new(sync.Map),
	}
}

// Lock locks the named lock for write access
func (nl *NamedLocker) Lock(name string) {
	mut, _ := nl.locks.LoadOrStore(name, new(sync.RWMutex))
	mut.(*sync.RWMutex).Lock()
}

// Unlock unlocks the named lock for write access, panics on non existence
func (nl *NamedLocker) Unlock(name string) {
	mut, _ := nl.locks.Load(name)
	mut.(*sync.RWMutex).Unlock()
}

// UnlockAndDelete unlocks the named lock for write access
// and removes it from memory, panics on non existence
func (nl *NamedLocker) UnlockAndDelete(name string) {
	mut, _ := nl.locks.Load(name)
	nl.locks.Delete(name)
	mut.(*sync.RWMutex).Unlock()
}

// RLock locks the named lock for read access
func (nl *NamedLocker) RLock(name string) {
	mut, _ := nl.locks.LoadOrStore(name, new(sync.RWMutex))
	mut.(*sync.RWMutex).RLock()
}

// RUnlock unlocks the named lock for read access, panics on non existence
func (nl *NamedLocker) RUnlock(name string) {
	mut, _ := nl.locks.Load(name)
	mut.(*sync.RWMutex).RUnlock()
}

// RUnlockAndDelete unlocks the named lock for read access and then deletes the lock, panics on non existence
func (nl *NamedLocker) RUnlockAndDelete(name string) {
	mut, _ := nl.locks.Load(name)
	mut.(*sync.RWMutex).RUnlock()
	nl.Delete(name)
}

// Delete deletes the named lock, panics on non existence
func (nl *NamedLocker) Delete(name string) {
	mut, _ := nl.locks.Load(name)
	mut.(*sync.RWMutex).Lock()
	defer mut.(*sync.RWMutex).Unlock()
	nl.locks.Delete(name)
}

package namedlocker

import "testing"

func TestMethods(t *testing.T) {
	lock := NewNamedLocker()
	lock.Lock("test")
	lock.Unlock("test")
	lock.RLock("test")
	lock.RUnlock("test")
	lock.Delete("test")

	lock.Lock("test")
	lock.UnlockAndDelete("test")
	lock.RLock("test")
	lock.RUnlockAndDelete("test")

}

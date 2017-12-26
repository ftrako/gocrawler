package test

import "sync"

// var queue sync.WaitGroup
var test_lock_locker sync.Mutex

func TestLock() {
	defer test_lock_locker.Unlock()

	test_lock_locker.Lock()
	// callWork()
}

func callWork() {
	defer test_lock_locker.Unlock()
	test_lock_locker.Lock()
}

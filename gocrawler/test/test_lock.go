package test

import (
	"fmt"
	"sync"
)

// var queue sync.WaitGroup
var test_lock_locker sync.Mutex

func TestLock() {
	defer test_lock_locker.Unlock()

	test_lock_locker.Lock()
	fmt.Println("TestLock")
	callWork()
}

func callWork() {
	defer test_lock_locker.Unlock()
	fmt.Println("callWork")
	test_lock_locker.Lock()
}

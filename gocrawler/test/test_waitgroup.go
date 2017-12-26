package test

import (
	"fmt"
	"sync"
	"time"
)

var queue sync.WaitGroup

var count = 0

func TestWait() {
	fmt.Println("started...")
	go newTest()
	queue.Add(1)

	queue.Wait()
	fmt.Println("finished!")
}

func newTest() {
	defer queue.Done()
	count++
	fmt.Println("newTest", time.Now(), count)
	if count > 100 {
		return
	}
	go newTest()
	queue.Add(1)
}

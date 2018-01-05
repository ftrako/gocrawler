package test

import (
	"fmt"
	"sync"
	"time"
)

var queue sync.WaitGroup

var count = 0

func TestWaitGroup() {
	fmt.Println("started...")
	go newTest()
	queue.Add(1)

	queue.Wait()
	fmt.Println("finished!")
}

func newTest() {
	defer func() {
		//queue.Done()

		if count > 100000 {
			queue.Done()
			return
		}
		go newTest()
	}()
	count++
	if (count % 10000) == 0 {
		fmt.Println("newTest", time.Now(), count)
	}
	time.Sleep(time.Millisecond * 10)
	//queue.Add(1)
}

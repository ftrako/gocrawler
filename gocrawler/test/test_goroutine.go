package test

import (
	"fmt"
	"strconv"
	"time"
)

var test_goroutine_ch = make(chan int, 1)
var test_goroutine_max = 10
var test_goroutine_count = 0

func TestGoRoutine() {
	for loop := 0; loop < 5; loop++ {
		//fmt.Println("toggle " + strconv.Itoa(loop))
		addWork("name"+strconv.Itoa(loop), test_goroutine_ch)
	}
}

func addWork(name string, ch chan int) {
	test_goroutine_count++
	if test_goroutine_count >= test_goroutine_max {
		return
	}
	fmt.Println("toggle " + name)
	ch <- 1
	go doWork(name, ch)
}

func doWork(name string, ch chan int) {
	defer func() {
		<-ch

		addWork("tmp212", ch)
	}()
	fmt.Println("name=" + name + ", len=" + strconv.Itoa(len(ch)) + ", cap=" + strconv.Itoa(cap(ch)))
	time.Sleep(time.Millisecond * 1000)
}

package test

import (
	"sync"
	"fmt"
	"time"
)

// var queue sync.WaitGroup
var test_error_locker sync.Mutex

func TestError() {
	defer func (){
		fmt.Println("TestError() end1",time.Now())
		if err := recover(); err != nil {
			fmt.Println(err,time.Now())
		}
		fmt.Println("TestError() end",time.Now())}()

	fmt.Println("TestError() start",time.Now())
	testError_fun1()
}

func testError_fun1() {
	defer fmt.Println("testError_fun1() end",time.Now())

	fmt.Println("testError_fun1() start",time.Now())
	//var a = 10
	//var b = 0
	//c := a /b
	//fmt.Println(c)
}

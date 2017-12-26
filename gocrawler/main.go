package main

import (
	"fmt"
	"gocrawler/crawler"
	"gocrawler/db"
	"time"
)

func main() {

	fmt.Println("start crawler...",time.Now())

	db.Open()
	crawler.SharedService().StartOneCrawler("appstore")
	//crawler.SharedService().StartOneCrawler("wandoujia")
	db.Close()

	// test.TestHelloWorld()

	// time.Sleep(time.Second * 50)

	fmt.Println("finished crawler! ", time.Now())
}

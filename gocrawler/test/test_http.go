package test

import (
	"net/http"
	"fmt"
	"time"
)

func TestHttp() {
	testHttp_doGet()
}

func testHttp_doGet() {
	fmt.Println("111", time.Now())
	url := "http://sw.bos.baidu.com/sw-search-sp/software/cf93beb4cb49a/BaiduMusic_10.1.10.0_setup.exe"
	_, err := http.Get(url)
	fmt.Println("222", time.Now())
	if err != nil {
		fmt.Println(err.Error())
	}
}

func testHttp_doPost() {

}

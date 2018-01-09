package test

import (
	"fmt"
	"gocrawler/util/httputil"
	"net/http"
	"time"
)

func TestHttp() {
	testHttp_doCustomGet()
}

func testHttp_doCustomGet() {
	fmt.Println("do 1")
	url := "http://www.java1234.com/a/kaiyuan/javaWeb/2013/0504/313.html"
	//url := "https://www.google.com"
	r, err := httputil.DoGetWithTimeout(url, time.Millisecond*20000)
	if err != nil {
		fmt.Println("err1", err)
		return
	}
	defer r.Body.Close()
	//c, err2 := ioutil.ReadAll(r.Body)
	//if err2 != nil {
	//	fmt.Println("err2", err2)
	//	return
	//}
	fmt.Println("content", r.ContentLength)
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

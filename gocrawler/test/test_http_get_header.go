package test

import (
	"fmt"
	"net/http"
)

func TestHttpGetHeader() {
	url := "https://download.jetbrains.8686c.com/go/goland-2017.3.dmg"
	url = "https://www.baidu.com"
	url = "http://www.51ape.com/tags"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("length", res.Header.Get("Content-Length"))
	fmt.Println(res.Header)
}

package main

import (
	"fmt"
	"time"
	"gocrawler/parser"
	"gocrawler/crawler"
	"os"
	"gocrawler/test"
	"strings"
	"strconv"
)

func main() {
	fmt.Println("start crawler...", time.Now())

	var parserType parser.ParserType = parser.ParserTypeWandoujia

	for _, value := range os.Args {
		if value == "" {
			continue
		}
		if strings.ToLower(value) == "test" {
			test.TestBackup()
			return
		}

		// 指定解析器
		if v, err := strconv.Atoi(value); err == nil {
			if v == 1 {
				parserType = parser.ParserTypeAppStore
			} else if v == 2 {
				parserType = parser.ParserTypeApe51
			}
		}
	}

	crawler.SharedService().StartOneCrawler(parserType)
	crawler.SharedService().Release()

	// time.Sleep(time.Second * 50)

	fmt.Println("finished crawler! ", time.Now())
}

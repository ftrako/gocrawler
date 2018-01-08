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

	var parserType parser.ParserType = parser.ParserTypeFile

	var restart = false

	for _, value := range os.Args {
		if value == "" {
			continue
		}
		if strings.ToLower(value) == "test" {
			test.TestCalc()
			return
		}

		if strings.ToLower(value) == "restart" {
			restart = true
		}

		// 指定解析器
		if v, err := strconv.Atoi(value); err == nil {
			switch v {
			case 0:
				parserType = parser.ParserTypeWandoujia
			case 1:
				parserType = parser.ParserTypeAppStore
			case 2:
				parserType = parser.ParserTypeApe51
			case 3:
				parserType = parser.ParserTypeAnn9
			case 4:
				parserType = parser.ParserTypeFile
			default:
			}
		}
	}

	if restart {
		crawler.SharedService().RestartOneCrawler(parserType)
	} else {
		crawler.SharedService().StartOneCrawler(parserType)
	}
	crawler.SharedService().Release()

	// time.Sleep(time.Second * 50)

	fmt.Println("finished crawler! ", time.Now())
}

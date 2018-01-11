package main

import (
	"fmt"
	"gocrawler/crawler"
	"gocrawler/parser"
	"gocrawler/test"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()
	fmt.Println("start crawler...", startTime)

	var parserType parser.ParserType = parser.ParserType_None

	var restart = false

	for _, value := range os.Args {
		if value == "" {
			continue
		}
		if strings.ToLower(value) == "test" {
			test.TestEnum()
			return
		}

		if strings.ToLower(value) == "restart" {
			restart = true
		}

		// 指定解析器
		if v, err := strconv.Atoi(value); err == nil {
			switch v {
			case 1:
				parserType = parser.ParserType_AndroidWandoujia
			case 2:
				parserType = parser.ParserType_AndroidAnzhi
			case 100:
				parserType = parser.ParserType_IosAppStore
			case 101:
				parserType = parser.ParserType_IosAnn9
			case 200:
				parserType = parser.ParserTypeApe51
			case 201:
				parserType = parser.ParserType_FileXuexi111
			case 202:
				parserType = parser.ParserType_FileDowncc
			case 203:
				parserType = parser.ParserType_FileGdajie
			case 204:
				parserType = parser.ParserType_FileJava1234
			case 205:
				parserType = parser.ParserType_FilePdfzj
			default:
			}
		}
	}

	started := false
	if restart {
		started = crawler.SharedService().RestartOneCrawler(parserType)
	} else {
		started = crawler.SharedService().StartOneCrawler(parserType)
	}
	crawler.SharedService().Release()

	if !started {
		fmt.Println("error: failed to start crawler")
	}

	// time.Sleep(time.Second * 50)

	endTime := time.Now()
	fmt.Println("finished crawler! ", endTime)
	fmt.Println("take times", endTime.Unix()-startTime.Unix(), "s")
}

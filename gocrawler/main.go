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

	var parserTypes = make([]parser.ParserType, 0, 100)

	var restart = false

	index := 0
	for _, value := range os.Args {
		index++
		if index == 1 {
			continue // 第一个是命令指令
		}
		if value == "" {
			continue
		}
		if strings.ToLower(value) == "test" {
			test.TestEnum()
			continue
		}

		if strings.ToLower(value) == "restart" {
			restart = true
			continue
		}

		var parserType parser.ParserType = parser.ParserType_None

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
		parserTypes = append(parserTypes, parserType)
	}

	if len(parserTypes) == 0 {
		parserTypes = append(parserTypes, parser.ParserType_None)
	}

	for _, v := range parserTypes {
		startCrawler(v, restart)
	}

	crawler.SharedService().Release()

	// time.Sleep(time.Second * 50)

	endTime := time.Now()
	fmt.Println("finished crawler! ", endTime)
	fmt.Println("take times", endTime.Unix()-startTime.Unix(), "s")
}

func startCrawler(parserType parser.ParserType, restart bool) {
	started := false
	if restart {
		started = crawler.SharedService().RestartOneCrawler(parserType)
	} else {
		started = crawler.SharedService().StartOneCrawler(parserType)
	}

	if !started {
		fmt.Println("error: failed to start crawler", parserType)
	}
}

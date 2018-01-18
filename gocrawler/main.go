package main

import (
	"fmt"
	"gocrawler/cmd"
	"gocrawler/crawler"
	"gocrawler/parser"
	"gocrawler/test"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	cmd.Cmd()

	startTime := time.Now()
	fmt.Println("start crawler...", startTime)

	var parserTypes = make([]parser.ParserType, 0, 100)

	var restart = false

	for _, value := range os.Args[1:] {
		if value == "" {
			continue
		}
		if strings.ToLower(value) == "test" {
			test.TestCmd()
			return
		}

		if strings.ToLower(value) == "restart" {
			restart = true
			continue
		}

		// 指定解析器
		if v, err := strconv.Atoi(value); err == nil {
			parserTypes = append(parserTypes, parser.ParserType(v))
		}
	}

	if len(parserTypes) == 0 {
		parserTypes = append(parserTypes, parser.ParserType_None)
		parserTypes = append(parserTypes, parser.ParserType_AndroidMi)
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

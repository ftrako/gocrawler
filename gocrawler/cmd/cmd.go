package cmd

import (
	"flag"
	"fmt"
)

func Cmd() {
	// -help or -h
	flag.Usage = func() {
		fmt.Println("使用方式参考：\n$ gocrawler test // 表示启动测试\n$ gocrawler 1 2 100 restart" +
			" // 表示启动爬虫，其中，restart参数表示重新爬，数字表示爬虫类型，目前支持如下类型：\n" +
			"1   // 安卓豌豆荚\n" +
			"2   // 安卓安智\n" +
			"100 // 苹果应用商店\n" +
			"101 // 苹果ann9爬的ios网站\n" +
			"200 // 51ape音乐网站\n" +
			"201 // Xuexi111文件网站\n" +
			"202 // Downcc文件网站\n" +
			"204 // Java1234文件网站\n" +
			"205 // FilePdfzj文件网站\n")
		flag.PrintDefaults()
	}
	flag.Parse()
}

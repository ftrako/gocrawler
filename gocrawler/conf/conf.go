package conf

import "runtime"

func GetMaxDepth() int {
	return 10
}

func GetDataPath() string {
	var ostype = runtime.GOOS
	if ostype == "windows" {
		return "data"
	}
	return "/data/go/gocrawler/data"
}

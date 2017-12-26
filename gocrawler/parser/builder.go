package parser

import (
	"gocrawler/data"
	"strings"
)

// NewParser 新建解析器
func NewParser(urlQueue *data.UrlQueue, parserId string) IParser {
	var p IParser
	parserId = strings.ToLower(parserId)
	if parserId == "appstore" {
		var p2 = new(AppStoreParser)
		p2.Os = "ios"
		p2.StoreId = "appstore"
		p2.StoreName = "苹果商店"
		p2.UrlQueue = urlQueue
		// p.StartUrl = "https://itunes.apple.com/cn/genre/ios/id36?mt=8"
		// p.StartUrl = "https://itunes.apple.com/cn/genre?id=36"
		p2.StartUrl = "https://itunes.apple.com/cn/app/%E5%A4%A9%E5%A4%A9%E7%88%B1%E6%B6%88%E9%99%A4/id654897098?mt=8"
		p = p2
	} else {
		var p2 = new(WandoujiaParser)
		p2.Os = "android"
		p2.StoreId = "wandoujia"
		p2.StoreName = "豌豆荚"
		p2.UrlQueue = urlQueue
		p2.StartUrl = "http://www.wandoujia.com/category/app"
		p = p2
	}
	return p
}

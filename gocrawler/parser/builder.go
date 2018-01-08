package parser

import (
	"gocrawler/db"
)

// NewParser 新建解析器
func NewParser(parserType ParserType) IParser {
	var p IParser
	switch parserType {
	case ParserTypeAppStore:
		p2 := new(AppStoreParser)
		p2.os = "ios"
		p2.storeId = "appstore"
		p2.storeName = "苹果商店"
		p2.id = p2.storeId
		p2.myDB = db.NewAppDB()
		p2.startUrl = "https://itunes.apple.com/cn/genre?id=36"
		p2.host = "https://itunes.apple.com"
		p = p2
	case ParserTypeWandoujia:
		var p2 = new(WandoujiaParser)
		p2.os = "android"
		p2.storeId = "wandoujia"
		p2.storeName = "豌豆荚"
		p2.id = p2.storeId
		p2.myDB = db.NewAppDB()
		p2.startUrl = "http://www.wandoujia.com/category/app"
		p2.host = "http://www.wandoujia.com"
		p = p2
	case ParserTypeApe51:
		var p2 = new(Ape51Parser)
		p2.myDB = db.NewMusicDB()
		p2.id = "ape51"
		p2.startUrl = "http://www.51ape.com/"
		p2.host = "http://www.51ape.com"
		p = p2
	case ParserTypeAnn9:
		p2 := new(Ann9Parser)
		p2.os = "ios"
		p2.storeId = "ann9"
		p2.storeName = "ann9"
		p2.id = p2.storeId
		p2.myDB = db.NewAppDB()
		p2.startUrl = "http://www.ann9.com/iphone"
		//p2.startUrl = "http://www.ann9.com/d33_11_1235504705_0.html"
		p2.host = "http://www.ann9.com"
		p = p2
	default:
	}

	return p
}

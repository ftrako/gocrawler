package parser

import (
	"github.com/PuerkitoBio/goquery"
)

type IParser interface {
	Parse(doc *goquery.Document) []string
	Filter(url string) bool // true表示满足爬虫过滤条件，允许爬
	GetStartUrl() string
	GetHost() string  // 获取域名，参考 http://www.example.com
	GetId() string
}

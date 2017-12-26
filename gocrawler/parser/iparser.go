package parser

import (
	"github.com/PuerkitoBio/goquery"
)

type IParser interface {
	Parse(doc *goquery.Document)
	GetStoreId() string
	GetOs() string
	GetStartUrl() string
}

package parser

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"gocrawler/db"
	"gocrawler/util/strutil"
)

type Ann9Parser struct {
	AppStoreParser
}

func (p *Ann9Parser) SetupData() {
	p.os = "ios"
	p.storeId = "ann9"
	p.storeName = "ann9"
	p.id = p.storeId
	p.myDB = db.NewAppDB()
	p.startUrl = "http://www.ann9.com/iphone"
}

func (p *Ann9Parser) Filter(url string) bool {
	if !strings.Contains(url, "ann9.com") {
		return false
	}
	if !p.BaseParser.Filter(url) {
		return false
	}
	return true
}

func (p *Ann9Parser) Parse(doc *goquery.Document) []string {
	urls := p.BaseParser.parseHref(doc)
	p.doParse(doc)
	return urls
}

func (p *Ann9Parser) doParse(doc *goquery.Document) {
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	p.parseApp(doc)
}

func (p *Ann9Parser) parseApp(doc *goquery.Document) {
	s := doc.Find("div.tel").Find("div.padbody2").Find("div.padappbody").Find("div.padsearch").Find("a.pkbtn").First()
	if s.Size() <= 0 {
		return
	}

	href, _ := s.Attr("href")
	if href == "" {
		return
	}

	index := strutil.LastIndex(href, "appid=") + len("appid=")
	id := strutil.SubString(href, index, strutil.Len(href))

	b := p.iosJsonParser.requestJsonByAppId(id)
	if b == nil {
		return
	}
	b.Os = p.os
	b.StoreId = p.storeId
	p.myDB.ReplaceApp(b)
}

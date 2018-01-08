package parser

import (
	"gocrawler/bean"
	"gocrawler/db"
	"gocrawler/util/strutil"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type AppStoreParser struct {
	BaseParser

	os        string // android or ios
	storeId   string
	storeName string
	myDB      *db.AppDB

	iosJsonParser *AppStoreParseJson
}

func (p *AppStoreParser) SetupData() {
	p.os = "ios"
	p.storeId = "appstore"
	p.storeName = "苹果商店"
	p.id = p.storeId
	p.myDB = db.NewAppDB()
	p.startUrl = "https://itunes.apple.com/cn/genre?id=36"
}

func (p *AppStoreParser) Release() {
	if p.myDB != nil {
		p.myDB.Close()
		p.myDB = nil
	}
}

func (p *AppStoreParser) Filter(url string) bool {
	if !strings.Contains(url, "apple.com") {
		return false
	}
	if !p.BaseParser.Filter(url) {
		return false
	}
	return true
}

func (p *AppStoreParser) Parse(doc *goquery.Document) []string {
	urls := p.BaseParser.parseHref(doc)
	p.doParse(doc)
	return urls
}

func (p *AppStoreParser) doParse(doc *goquery.Document) {
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	// 爬分类
	p.parseCategory(doc)

	// 爬应用
	p.parseApp(doc)
}

func (p *AppStoreParser) parseCategory(doc *goquery.Document) {
	if doc == nil {
		return
	}

	var level0Name string

	// 总共三级

	// 第一级
	doc.Find("div#content").Find("div#media-type-nav.nav").Find("ul.list").Each(func(i int, s *goquery.Selection) {
		s.Find("li").Find("a").Each(func(i int, s2 *goquery.Selection) {
			href, _ := s2.Attr("href")
			if href == "" {
				return
			}
			var c bean.CategoryBean
			c.Name = s2.Text()
			c.Cid = p.parseCid(href)
			c.StoreId = p.storeId
			c.StoreName = p.storeName
			p.myDB.ReplaceCategory(&c)
		})

		// 当前selected项
		s2 := s.Find("li.selected").First()
		level0Name = s2.Text()
	})

	// 第二级
	doc.Find("div#content").Find("div.main.nav#genre-nav").Find("div").Find("ul.list.column").Find("li").Each(func(i int, s *goquery.Selection) {
		var level1Name string
		s2 := s.Find("a.top-level-genre").First()
		level1Name = s2.Text()
		href, _ := s2.Attr("href")
		if href == "" {
			return
		}

		var c bean.CategoryBean
		c.Cid = p.parseCid(href)
		c.Name = s2.Text()
		c.StoreId = p.storeId
		c.StoreName = p.storeName
		c.SuperName = level0Name
		p.myDB.ReplaceCategory(&c)

		// 第三级
		s.Find("ul").Find("li").Find("a").Each(func(i int, s2 *goquery.Selection) {
			href, _ := s2.Attr("href")
			if href == "" {
				return
			}

			var c2 bean.CategoryBean
			c2.Cid = p.parseCid(href)
			c2.Name = s2.Text()
			c2.StoreId = p.storeId
			c2.StoreName = p.storeName
			c2.SuperName = level1Name
			p.myDB.ReplaceCategory(&c2)
		})

	})
}

func (p *AppStoreParser) parseApp(doc *goquery.Document) {
	s := doc.Find("div#left-stack").Find("a").First()
	text, _ := s.Attr("href")
	if text == "" {
		return
	}

	start := strings.LastIndex(text, "/") + 3
	end := strings.LastIndex(text, "?")
	text = strutil.SubString(text, start, end)

	b := p.iosJsonParser.requestJsonByAppId(text)
	if b == nil {
		return
	}

	b.Os = p.os
	b.StoreId = p.storeId

	p.myDB.ReplaceApp(b)
}

func (p *AppStoreParser) parseCid(text string) string {
	index := strutil.LastIndex(text, "/id") + strutil.Len("/id")
	cid := strutil.SubString(text, index, strutil.Len(text))
	index = strutil.Index(cid, "?")
	if index >= 0 { // 有的链接带有？
		cid = strutil.SubString(cid, 0, index)
	}
	cid = strings.Replace(cid, " ", "", -1)
	return cid
}

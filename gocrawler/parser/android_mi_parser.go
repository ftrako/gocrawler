package parser

import (
	"github.com/PuerkitoBio/goquery"
	"gocrawler/bean"
	"gocrawler/util/strutil"
	"strings"
)

type AndroidMiParser struct {
	BaseAppParser
}

func (p *AndroidMiParser) SetupData() {
	p.BaseAppParser.SetupData()
	p.os = "android"
	p.storeId = "mi"
	p.storeName = "小米商店"
	p.id = p.storeId
	p.startUrl = "http://app.mi.com/category/1"
	p.startUrl = "http://app.mi.com/details?id=com.jianlv.chufaba"
}

func (p *AndroidMiParser) Filter(url string) bool {
	if !strings.Contains(url, "mi.com/category") &&
		!strings.Contains(url, "mi.com/details") {
		return false
	}
	if !p.BaseParser.Filter(url) {
		return false
	}
	return true
}

func (p *AndroidMiParser) Parse(doc *goquery.Document) []string {
	urls := p.BaseParser.Parse(doc)

	// 部分url需要处理，不然会跳转到自动下载apk
	size := len(urls)
	for loop := 0; loop < size; loop++ {
		url := urls[loop]
		if strings.HasSuffix(url, "/binding") {
			url = strutil.SubString(url, 0, strutil.Len(url)-len("/binding"))
			urls[loop] = url
		} else if strings.Contains(url, "/binding?") {
			index := strutil.Index(url, "/binding?")
			url = strutil.SubString(url, 0, index)
			urls[loop] = url
		} else if strings.Contains(url, "/download?") {
			index := strutil.Index(url, "/download?")
			url = strutil.SubString(url, 0, index)
			urls[loop] = url
		}
	}
	p.doParse(doc)
	return urls
}

func (p *AndroidMiParser) doParse(doc *goquery.Document) {
	defer func() {
		if err := recover(); err != nil {
		}
	}()

	// 爬分类
	//p.parseCategory(doc)

	// 爬应用
	p.parseApp(doc)
}

func (p *AndroidMiParser) parseCategory(doc *goquery.Document) {
	if doc == nil {
		return
	}

	doc.Find("li.app-tag-wrap").Find("a.app-tag").Find("span").Each(func(i int, s *goquery.Selection) {
		// 插入第一层
		var b bean.CategoryBean
		b.Name = s.Text()
		p.myDB.ReplaceCategory(&b)
		p.parseSubCategory(doc, "li.app-tag-wrap", s.Text())
	})

	doc.Find("li.game-tag-wrap").Find("a.game-tag").Find("span").Each(func(i int, s *goquery.Selection) {
		// 插入第一层
		var b bean.CategoryBean
		b.Name = s.Text()
		p.myDB.ReplaceCategory(&b)
		p.parseSubCategory(doc, "li.game-tag-wrap", s.Text())
	})
}

func (p *AndroidMiParser) parseSubCategory(doc *goquery.Document, basequery string, basename string) {
	if doc == nil {
		return
	}
	doc.Find(basequery).Find("li.parent-cate").Each(func(i int, subs *goquery.Selection) {
		subs.Find("a.cate-link").Each(func(j int, subss *goquery.Selection) {
			subs.Find("li.child-cate").Each(func(j int, subsss *goquery.Selection) {
				subsss.Find("a").Each(func(k int, subssss *goquery.Selection) {
					var b bean.CategoryBean
					b.SuperName = subss.Text()
					b.StoreId = p.storeId
					b.StoreName = p.storeName
					b.Name = subssss.Text()
					p.myDB.ReplaceCategory(&b) // 插入分类
				})
			})
		})
	})
}

func (p *AndroidMiParser) parseApp(doc *goquery.Document) {
	var b bean.AppBean
	b.Os = p.os
	b.StoreId = p.storeId

	// appid
	url := doc.Url.String()
	start := strutil.LastIndex(url, "details?id=") + strutil.Len("details?id=")
	end := strutil.Len(url)
	b.AppId = strutil.SubString(url, start, end)
	if b.AppId == "" {
		return
	}

	b.Name = doc.Find("div.intro-titles").Find("h3").First().Text()
	if b.Name == "" {
		return
	}

	b.Category = doc.Find("div.intro-titles").Find("p.action").First().Text()
	start = 3
	end = strutil.Index(b.Category, "|")
	b.Category = strutil.SubString(b.Category, start, end)
	b.Category = ";" + b.Category + ";"

	index := 0
	doc.Find("div.look-detail").Find("div.details").Find("li").Each(func(j int, s *goquery.Selection) {
		index++
		if index == 2 {
			b.Size = s.Text()
		} else if index == 4 {
			b.Version = s.Text()
		} else if index == 6 {
			b.UpdateTime = s.Text()
		}
	})

	p.myDB.ReplaceApp(&b)
}

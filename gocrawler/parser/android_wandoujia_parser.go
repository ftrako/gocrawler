package parser

import (
	"gocrawler/bean"
	"gocrawler/util/strutil"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type AndroidWandoujiaParser struct {
	BaseAppParser
}

func (p *AndroidWandoujiaParser) SetupData() {
	p.BaseAppParser.SetupData()
	p.os = "android"
	p.storeId = "wandoujia"
	p.storeName = "豌豆荚"
	p.id = p.storeId
	p.startUrl = "http://www.wandoujia.com/apps"
}

func (p *AndroidWandoujiaParser) Filter(url string) bool {
	if !strings.Contains(url, "wandoujia.com/apps") &&
		!strings.Contains(url, "wandoujia.com/category") &&
		!strings.Contains(url, "wandoujia.com/top") &&
		!strings.Contains(url, "wandoujia.com/essential") &&
		!strings.Contains(url, "wandoujia.com/special") {
		return false
	}
	if !p.BaseParser.Filter(url) {
		return false
	}
	return true
}

func (p *AndroidWandoujiaParser) Parse(doc *goquery.Document) []string {
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

func (p *AndroidWandoujiaParser) doParse(doc *goquery.Document) {
	defer func() {
		if err := recover(); err != nil {
		}
	}()

	// 爬分类
	p.parseCategory(doc)

	// 爬应用
	p.parseApp(doc)
}

func (p *AndroidWandoujiaParser) parseCategory(doc *goquery.Document) {
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

func (p *AndroidWandoujiaParser) parseSubCategory(doc *goquery.Document, basequery string, basename string) {
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

func (p *AndroidWandoujiaParser) parseApp(doc *goquery.Document) {
	var b bean.AppBean
	b.Os = p.os
	b.StoreId = p.storeId
	doc.Find("div.detail-wrap").Find("div.detail-top.clearfix").Find("div.app-info").Find("a[data-app-id]").Each(func(j int, s *goquery.Selection) {
		var text string
		text, _ = s.Attr("data-app-name")
		if text == "" {
			return // 没有名称
		}
		b.Name = text
		text, _ = s.Attr("data-app-pname")
		if text == "" {
			return // 没有appid
		}
		b.AppId = text
		text, _ = s.Attr("data-app-vname")
		b.Version = text
		text, _ = s.Attr("data-install")
		b.InstallCount = text
	})
	doc.Find("div.detail-wrap").Find("div.infos").Find("dl.infos-list").Each(func(j int, s *goquery.Selection) {
		s.Find("meta[itemprop][content]").Each(func(j int, ss *goquery.Selection) {
			text, _ := ss.Attr("content")
			b.Size = text
		})
		var categories string
		s.Find("dd.tag-box").Find("a").Each(func(j int, ss *goquery.Selection) {
			categories += ss.Text() + ";"
		})
		b.Category = categories
		s.Find("time[datetime]").Each(func(j int, ss *goquery.Selection) {
			b.UpdateTime = ss.Text()
		})
		s.Find("span.dev-sites").Each(func(j int, ss *goquery.Selection) {
			b.Vender = ss.Text()
		})
		s.Find("dd.perms[itemprop=operatingSystems]").Each(func(j int, ss *goquery.Selection) {
			text := ss.Text()
			text = strings.Trim(text, "\n")
			index := strings.Index(text, "\n")
			text = strutil.SubString(text, 0, index)
			text = strings.TrimSpace(text)
			b.MinVersion = text
		})
	})

	p.myDB.ReplaceApp(&b)
}

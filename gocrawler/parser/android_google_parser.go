package parser

import (
	"gocrawler/bean"
	"gocrawler/util/strutil"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type AndroidGoogleParser struct {
	BaseAppParser
}

func (p *AndroidGoogleParser) SetupData() {
	p.BaseAppParser.SetupData()
	p.os = "android"
	p.storeId = "google"
	p.storeName = "谷歌"
	p.id = p.storeId
	p.startUrl = "https://play.google.com/store/apps"
	//p.startUrl = "https://play.google.com/store/apps/details?id=com.tencent.mobileqq"
}

func (p *AndroidGoogleParser) Filter(url string) bool {
	//if !strings.Contains(url, "wandoujia.com/apps") &&
	//	!strings.Contains(url, "wandoujia.com/category") &&
	//	!strings.Contains(url, "wandoujia.com/top") &&
	//	!strings.Contains(url, "wandoujia.com/essential") &&
	//	!strings.Contains(url, "wandoujia.com/special") {
	//	return false
	//}
	if !p.BaseParser.Filter(url) {
		return false
	}
	return true
}

func (p *AndroidGoogleParser) Parse(doc *goquery.Document) []string {
	urls := p.BaseParser.Parse(doc)

	// 部分url需要处理，不然会跳转到自动下载apk
	//size := len(urls)
	//for loop := 0; loop < size; loop++ {
	//	url := urls[loop]
	//	if strings.HasSuffix(url, "/binding") {
	//		url = strutil.SubString(url, 0, strutil.Len(url)-len("/binding"))
	//		urls[loop] = url
	//	} else if strings.Contains(url, "/binding?") {
	//		index := strutil.Index(url, "/binding?")
	//		url = strutil.SubString(url, 0, index)
	//		urls[loop] = url
	//	} else if strings.Contains(url, "/download?") {
	//		index := strutil.Index(url, "/download?")
	//		url = strutil.SubString(url, 0, index)
	//		urls[loop] = url
	//	}
	//}
	p.doParse(doc)
	return urls
}

func (p *AndroidGoogleParser) doParse(doc *goquery.Document) {
	defer func() {
		if err := recover(); err != nil {
		}
	}()

	// 爬分类
	//p.parseCategory(doc)

	// 爬应用
	p.parseApp(doc)
}

func (p *AndroidGoogleParser) parseCategory(doc *goquery.Document) {
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

func (p *AndroidGoogleParser) parseSubCategory(doc *goquery.Document, basequery string, basename string) {
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

func (p *AndroidGoogleParser) parseApp(doc *goquery.Document) {
	var b bean.AppBean
	b.Os = p.os
	b.StoreId = p.storeId

	url := doc.Url.String()
	start := strutil.LastIndex(url, "?id=") + strutil.Len("?id=")
	end := strutil.Len(url)
	b.AppId = strutil.SubString(url, start, end)
	if b.AppId == "" {
		return
	}

	b.Name = doc.Find("h1.document-title").Find("div.id-app-title").First().Text()
	if b.Name == "" {
		return
	}

	b.Category = doc.Find("div.author").Find("a.document-subtitle.category").Find("span").First().Text()
	b.Category = ";" + b.Category + ";"

	index := 0
	doc.Find("div.details-section.metadata").Find("div.details-section-contents").Find("div.meta-info").Each(func(j int, s *goquery.Selection) {
		index++
		if index == 1 { // update time
			b.UpdateTime = s.Find("div.content").First().Text()
		} else if index == 2 { // install count
			b.InstallCount = s.Find("div.content").First().Text()
		} else if index == 3 { // version
			b.Version = s.Find("div.content").First().Text()
		} else if index == 4 { // min version
			b.MinVersion = s.Find("div.content").First().Text()
		} else if index == 9 { // vender
			b.Vender = s.Find("div.content").First().Text()
		}
	})

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

package parser

import (
	"gocrawler/bean"
	"gocrawler/util/strutil"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type AndroidAnzhiParser struct {
	BaseAppParser
}

func (p *AndroidAnzhiParser) SetupData() {
	p.BaseAppParser.SetupData()
	p.os = "android"
	p.storeId = "anzhi"
	p.storeName = "安智"
	p.id = p.storeId
	//p.startUrl = "http://www.anzhi.com/applist.html"
	p.startUrl = "http://www.anzhi.com/pkg/7093_com.fyhb.app.html"
}

func (p *AndroidAnzhiParser) Filter(url string) bool {
	if !p.BaseParser.Filter(url) {
		return false
	}

	if !strings.Contains(url, "anzhi.com/applist.html") &&
		!strings.Contains(url, "anzhi.com/gamelist.html") &&
		!strings.Contains(url, "anzhi.com/pkg") &&
		!strings.Contains(url, "anzhi.com/sort") {
		return false
	}

	return true
}

func (p *AndroidAnzhiParser) Parse(doc *goquery.Document) []string {
	urls := p.BaseParser.Parse(doc)

	//// 部分url需要处理，不然会跳转到自动下载apk
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

func (p *AndroidAnzhiParser) doParse(doc *goquery.Document) {
	defer func() {
		if err := recover(); err != nil {
		}
	}()

	// 爬应用
	p.parseApp(doc)
}

func (p *AndroidAnzhiParser) parseApp(doc *goquery.Document) {
	var b bean.AppBean
	b.Os = p.os
	b.StoreId = p.storeId

	b.AppId, _ = doc.Find("head").Find("meta[property=\"og:soft:url\"]").First().Attr("content")
	start := strutil.Index(b.AppId, "_") + 1
	end := strutil.LastIndex(b.AppId, ".html")
	b.AppId = strutil.SubString(b.AppId, start, end)
	if b.AppId == "" {
		return
	}

	b.Name = doc.Find("div.detail_description").Find("div.detail_line").Find("h3").First().Text()
	if b.Name == "" {
		return
	}
	b.Version = doc.Find("div.detail_description").Find("div.detail_line").Find("span.app_detail_version").First().Text()
	b.Version = strutil.SubString(b.Version, 1, strutil.Len(b.Version)-1)

	index := 0
	doc.Find("div.detail_description").Find("ul#detail_line_ul").Find("li").Each(func(j int, s *goquery.Selection) {
		index++
		if index == 1 { // category
			b.Category = s.Text()
			b.Category = ";" + strutil.SubString(b.Category, 3, strutil.Len(b.Category)) + ";"
		} else if index == 2 { // install count
			b.InstallCount = s.Find("span").First().Text()
			b.InstallCount = strutil.SubString(b.InstallCount, 3, strutil.Len(b.InstallCount))
		} else if index == 3 { // update date
			b.UpdateTime = s.Text()
			b.UpdateTime = strutil.SubString(b.UpdateTime, 3, strutil.Len(b.UpdateTime))
		} else if index == 4 { // size
			b.Size = s.Find("span").First().Text()
			b.Size = strutil.SubString(b.Size, 3, strutil.Len(b.Size))
		} else if index == 5 { //min version
			b.MinVersion = s.Text()
			b.MinVersion = strutil.SubString(b.MinVersion, 3, strutil.Len(b.MinVersion))
		} else if index == 7 { // vender
			b.Vender = s.Text()
			b.Vender = strutil.SubString(b.Vender, 3, strutil.Len(b.Vender))
		}
	})

	p.myDB.ReplaceApp(&b)
}

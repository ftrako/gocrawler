package parser

import (
	"fmt"
	"gocrawler/bean"
	"gocrawler/util/strutil"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type AndroidQqParser struct {
	BaseAppParser
}

func (p *AndroidQqParser) SetupData() {
	p.BaseAppParser.SetupData()
	p.os = "android"
	p.storeId = "qq"
	p.storeName = "应用宝"
	p.id = p.storeId
	p.startUrl = "http://sj.qq.com/myapp/category.htm?orgame=1"
}

func (p *AndroidQqParser) Filter(url string) bool {
	if !strings.Contains(url, "qq.com/myapp/category.htm") &&
		!strings.Contains(url, "qq.com/myapp/detail.htm") {
		return false
	}
	if !p.BaseParser.Filter(url) {
		return false
	}
	return true
}

func (p *AndroidQqParser) Parse(doc *goquery.Document) []string {
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

func (p *AndroidQqParser) doParse(doc *goquery.Document) {
	defer func() {
		if err := recover(); err != nil {
		}
	}()

	// 爬分类
	//p.parseCategory(doc)

	// 爬应用
	p.parseApp(doc)
}

func (p *AndroidQqParser) parseCategory(doc *goquery.Document) {
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

func (p *AndroidQqParser) parseSubCategory(doc *goquery.Document, basequery string, basename string) {
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

func (p *AndroidQqParser) parseApp(doc *goquery.Document) {
	var b bean.AppBean
	b.Os = p.os
	b.StoreId = p.storeId

	// appid
	url := doc.Url.String()
	start := strutil.LastIndex(url, "?apkName=") + strutil.Len("?apkName=")
	end := strutil.Len(url)
	b.AppId = strutil.SubString(url, start, end)
	if b.AppId == "" {
		return
	}

	b.Name = doc.Find("div.det-main-container").Find("div.det-ins-data").Find("div.det-name").Find("div.det-name-int").First().Text()
	if b.Name == "" {
		return
	}

	b.Category = doc.Find("div.det-main-container").Find("div.det-ins-data").Find("div.det-type-box").Find("a").First().Text()
	b.Category = ";" + b.Category + ";"

	b.InstallCount = doc.Find("div.det-main-container").Find("div.det-ins-data").Find("div.det-insnum-line").Find("div.det-ins-num").First().Text()
	b.InstallCount = strutil.SubString(b.InstallCount, 0, strutil.Len(b.InstallCount)-2)

	b.Size = doc.Find("div.det-main-container").Find("div.det-ins-data").Find("div.det-insnum-line").Find("div.det-size").First().Text()

	index := 0
	doc.Find("div.det-othinfo-container").Find("div.det-othinfo-data").Each(func(j int, s *goquery.Selection) {
		index++
		if index == 1 { // version
			b.Version = s.Text()
		} else if index == 2 { // update time
			b.UpdateTime, _ = s.Attr("data-apkpublishtime")
			if sec, err := strconv.Atoi(b.UpdateTime); err == nil {
				t := time.Unix(int64(sec), 0)
				b.UpdateTime = fmt.Sprintf("%04d年%02d月%02d日", t.Year(), t.Month(), t.Day())
			}
		} else if index == 3 { // vender
			b.Vender = s.Text()
		}
	})

	p.myDB.ReplaceApp(&b)
}

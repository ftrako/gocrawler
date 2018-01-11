package parser

import (
	"github.com/PuerkitoBio/goquery"
	"gocrawler/bean"
	"gocrawler/util/strutil"
	"strings"
)

// xuexi111.com

type Xuexi111FileParser struct {
	BaseFileParser
}

func (p *Xuexi111FileParser) SetupData() {
	p.BaseFileParser.SetupData()
	p.startUrl = "http://www.xuexi111.com/book/"
}

func (p *Xuexi111FileParser) Filter(url string) bool {
	if !p.BaseParser.Filter(url) {
		return false
	}

	if !strings.Contains(url, "xuexi111.com/book/") {
		return false
	}

	return true
}

func (p *Xuexi111FileParser) Parse(doc *goquery.Document) []string {
	urls := p.BaseFileParser.Parse(doc)
	p.doParse(doc)
	return urls
}

func (p *Xuexi111FileParser) doParse(doc *goquery.Document) {
	defer func() {
		if err := recover(); err != nil {
		}
	}()

	// 获取日期
	index := 0
	var updateDate string
	doc.Find("div.txt_info").Find("div.cont").Find("div.cont_l").Find("ul").Find("li").Each(func(i int, s *goquery.Selection) {
		index++
		if index == 7 {
			updateDate = s.Find("span").First().Text()
		}
	})

	// 因有多个下载链接，所以只抓下载链接
	doc.Find("div#download.download-url").Find("table.download-table#download-table").Find("tr").Each(func(i int, s *goquery.Selection) {
		b := new(bean.FileBean)
		a := s.Find("a").First()
		b.Download, _ = a.Attr("href")
		if b.Download == "" {
			return
		}

		b.Name = a.Text()
		if b.Name == "" {
			return
		}
		b.Suffix = strings.ToLower(strutil.Suffix(b.Name))
		if b.Suffix != "pdf" && b.Suffix != "epub" { // 只抓pdf和epub
			return
		}
		b.Size = s.Find("td[align=\"center\"]").First().Text()

		if strutil.Len(b.Name) > 200 {
			b.Name = strutil.SubString(b.Name, 0, 200)
		}

		b.UpdateDate = updateDate
		b.Url = doc.Url.String()
		b.Type = "book"
		p.myDB.ReplaceFile(b)
	})
}

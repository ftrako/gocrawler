package parser

import (
	"github.com/PuerkitoBio/goquery"
	"gocrawler/bean"
	"gocrawler/util/strutil"
	"strings"
)

// 被防爬了

type GdajieFileParser struct {
	BaseFileParser
}

func (p *GdajieFileParser) SetupData() {
	p.BaseFileParser.SetupData()
	p.startUrl = "http://verycd.gdajie.com/book/page"
}

func (p *GdajieFileParser) Filter(url string) bool {
	if !p.BaseParser.Filter(url) {
		return false
	}

	if !strings.Contains(url, "gdajie.com") {
		return false
	}

	return true
}

func (p *GdajieFileParser) Parse(doc *goquery.Document) []string {
	urls := p.BaseFileParser.Parse(doc)
	p.doParse(doc)
	return urls
}

func (p *GdajieFileParser) doParse(doc *goquery.Document) {
	defer func() {
		if err := recover(); err != nil {
		}
	}()

	s := doc.Find("div.body").Find("div.main").Find("center").Find("div#detail").Find("table").Find("td").Find("div").Find("center").Find("a").First()
	if s.Size() <= 0 {
		return
	}

	name := s.Text()
	if name == "" {
		return
	}

	if strutil.Len(name) > 200 {
		name = strutil.SubString(name, 0, 200)
	}

	link, _ := s.Attr("href")
	if len(link) > 280 {
		return
	}
	if link == "" {
		return
	}

	b := new(bean.FileBean)
	b.Name = name
	b.Suffix = strutil.Suffix(name)
	b.Url = doc.Url.String()
	b.Download = link
	b.Type = "book"
	p.myDB.ReplaceFile(b)
}

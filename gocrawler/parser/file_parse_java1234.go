package parser

import (
	"github.com/PuerkitoBio/goquery"
	"gocrawler/bean"
	"gocrawler/util/strutil"
	"strings"
)

type Java1234FileParser struct {
	BaseFileParser
}

func (p *Java1234FileParser) SetupData() {
	p.BaseFileParser.SetupData()
	p.startUrl = "http://www.java1234.com"
}

func (p *Java1234FileParser) Filter(url string) bool {
	if !p.BaseParser.Filter(url) {
		return false
	}

	if !strings.Contains(url, "java1234.com") {
		return false
	}

	return true
}

func (p *Java1234FileParser) Parse(doc *goquery.Document) []string {
	urls := p.BaseFileParser.Parse(doc)
	p.doParse(doc)
	return urls
}

func (p *Java1234FileParser) doParse(doc *goquery.Document) {
	defer func() {
		if err := recover(); err != nil {
		}
	}()

	s := doc.Find("head").Find("title").First()
	name := s.Text()
	index := strutil.Index(name, " ")
	if index > 0 {
		name = strutil.SubString(name, 0, index)
	}
	if strutil.Len(name) > 100 {
		name = strutil.SubString(name, 0, 100)
	}
	if !strings.Contains(strings.ToLower(name), "pd") { // 标题不含pd未非电子书，有时pdf被截成pd
		return
	}
	s = doc.Find("div.content").Find("table").Find("td").Find("span").Find("strong").Find("a").First()
	if s.Size() <= 0 {
		return
	}
	link, _ := s.Attr("href")
	if len(link) > 260 {
		return
	}
	s = doc.Find("div.content").Find("table").Find("td").Find("span").Find("strong").Find("span").First()
	if s.Size() <= 0 {
		return
	}
	pwd := s.Text()

	if strings.Contains(pwd, "主要内容") || len(pwd) > 8 { // 异常数据
		return
	}

	if link == "" {
		return
	}
	b := new(bean.FileBean)
	b.Name = name
	b.Url = doc.Url.String()
	b.Download = link
	b.Pwd = pwd
	b.Type = "book"
	b.Suffix = "pdf"
	p.myDB.ReplaceFile(b)
}

package parser

import (
	"github.com/PuerkitoBio/goquery"
	"gocrawler/bean"
	"gocrawler/db"
	"gocrawler/util/strutil"
	"strings"
)

// downcc.com

type DownccFileParser struct {
	BaseFileParser
}

func (p *DownccFileParser) SetupData() {
	p.BaseFileParser.SetupData()
	p.startUrl = "http://www.downcc.com/soft/list_181_1.html"
}

func (p *DownccFileParser) Filter(url string) bool {
	if !p.BaseParser.Filter(url) {
		return false
	}

	if !strings.Contains(url, "downcc.com/soft/list_181") {
		return false
	}

	return true
}

func (p *DownccFileParser) Parse(doc *goquery.Document) []string {
	urls := p.BaseFileParser.Parse(doc)
	p.doParse(doc)
	return urls
}

func (p *DownccFileParser) doParse(doc *goquery.Document) {
	defer func() {
		if err := recover(); err != nil {
		}
	}()

	doc.Find("ul#li-change-color").Find("li").Each(func(i int, s *goquery.Selection) {
		b := new(bean.FileBean)
		a := s.Find("h3.soft-ht1").Find("a").First()
		b.Name = strutil.Gbk2Utf8(a.Text())
		if b.Name == "" {
			return
		}
		link, _ := a.Attr("href")
		b.Url = "http://www.downcc.com" + link
		b.Size = strutil.Gbk2Utf8(s.Find("p").Find("span.mg-r10").First().Text())
		b.Size = strutil.SubString(b.Size, 3, strutil.Len(b.Size))
		b.UpdateDate = strutil.Gbk2Utf8(s.Find("p.soft-update").Text())
		b.UpdateDate = strutil.SubString(b.UpdateDate, 5, strutil.Len(b.UpdateDate))

		b.Type = "book"
		p.myDB.ReplaceFile(b)
	})
}

func (p *DownccFileParser) FileParse_old(doc *goquery.Document, db *db.FileDB) {
	index := 0
	b := new(bean.FileBean)
	s := doc.Find("p#topNavC").Find("strong").First()
	b.Name = strutil.Gbk2Utf8(s.Text())
	if b.Name == "" {
		return
	}

	link := doc.Find("ul.ul_Address").Find("script").First().Text()
	start := len(" _downInfo = {Address:\"")
	end := strutil.Index(link, "\",")
	link = strutil.SubString(link, start, end)
	b.Download = "http://js.downcc.com" + link

	isBook := false

	doc.Find("section.layout.clear").Find("section.soft-details").Find("article").Find("ul").Find("li").Each(func(i int, s *goquery.Selection) {
		index++
		if index == 1 {
			b.Size = strutil.Gbk2Utf8(s.Text())
			b.Size = strutil.SubString(b.Size, 5, strutil.Len(b.Size))
		} else if index == 4 {
			t := strutil.Gbk2Utf8(s.Text())
			if strings.Contains(t, "图书") {
				isBook = true
			}
		} else if index == 5 {
			b.UpdateDate = strutil.Gbk2Utf8(s.Text())
			b.UpdateDate = strutil.SubString(b.UpdateDate, 5, strutil.Len(b.UpdateDate))
		}
	})

	if !isBook { // 只爬图书
		return
	}

	if strutil.Len(b.Name) > 200 {
		b.Name = strutil.SubString(b.Name, 0, 200)
	}

	b.Suffix = strings.ToLower(strutil.Suffix(b.Download))
	if b.Suffix == "com" {
		return // 假电子书
	}
	b.Url = doc.Url.String()
	b.Type = "book"
	db.ReplaceFile(b)
}

package parser

import (
	"github.com/PuerkitoBio/goquery"
	"gocrawler/bean"
	"gocrawler/db"
	"gocrawler/util/strutil"
	"strings"
)

type Java1234FileParser struct {
}

func (p *Java1234FileParser) StartUrl() string {
	return "http://www.java1234.com"
}

func (p *Java1234FileParser) FileFilter(url string) bool {
	if strings.Contains(url, "java1234.com") {
		return true
	} else {
		return false
	}
}

func (p *Java1234FileParser) FileParse(doc *goquery.Document, db *db.FileDB) {
	s := doc.Find("head").Find("title").First()
	name := s.Text()
	index := strutil.Index(name, " ")
	if index > 0 {
		name = strutil.SubString(name, 0, index)
	}
	if strutil.Len(name) > 100 {
		name = strutil.SubString(name, 0, 100)
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
	db.ReplaceFile(b)
}

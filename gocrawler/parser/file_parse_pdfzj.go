package parser

import (
	"github.com/PuerkitoBio/goquery"
	"gocrawler/bean"
	"gocrawler/util/strutil"
	"strings"
)

// pdfzj.cn

type PdfzjFileParser struct {
	BaseFileParser
}

func (p *PdfzjFileParser) SetupData() {
	p.BaseFileParser.SetupData()
	p.startUrl = "http://pdfzj.cn/"
}

func (p *PdfzjFileParser) Filter(url string) bool {
	if !p.BaseParser.Filter(url) {
		return false
	}

	if !strings.Contains(url, "pdfzj.cn") {
		return false
	}

	return true
}

func (p *PdfzjFileParser) Parse(doc *goquery.Document) []string {
	urls := p.BaseFileParser.Parse(doc)
	p.doParse(doc)
	return urls
}

func (p *PdfzjFileParser) doParse(doc *goquery.Document) {
	defer func() {
		if err := recover(); err != nil {
		}
	}()

	index := 0
	b := new(bean.FileBean)

	isBook := false

	doc.Find("div.list-list").Find("div.bookes-specifics").Find("div.books-spe").Find("div.table_tr").Each(func(i int, s *goquery.Selection) {
		index++
		s2 := s.Find("div.table_td2").First()
		if index == 1 { // name
			s3 := s2.Find("font").First()
			b.Name = s3.Text()
		} else if index == 2 { // author
			s3 := s2.Find("strong").First()
			b.Author = s3.Text()
		} else if index == 3 { // size
			s3 := s2.Find("strong").First()
			b.Size = s3.Text()
		} else if index == 4 { // suffix
			s3 := s2.Find("strong").First()
			b.Suffix = strings.ToLower(s3.Text())
			if b.Suffix == "pdf" {
				isBook = true
			}
		} else if index == 6 { // unzip_pwd
			s3 := s2.Find("font").First()
			b.UnzipPwd = s3.Text()
		} else if index == 7 { // update_date
			s3 := s2.Find("strong").First()
			b.UpdateDate = s3.Text()
		}
	})

	if b.Name == "" || !isBook { // 只爬pdf
		return
	}

	s := doc.Find("h1").Find("span").First()
	b.Pwd = s.Text()

	s = doc.Find("div.download_area#download_area__").Find("div.table_tr.down_area").Find("div.table_td.g-left").Find("strong").Find("a").First()
	b.Download, _ = s.Attr("href")
	if b.Download == "" {
		return
	}

	if strutil.Len(b.Name) > 200 {
		b.Name = strutil.SubString(b.Name, 0, 200)
	}

	b.Url = doc.Url.String()
	b.Type = "book"
	p.myDB.ReplaceFile(b)
}

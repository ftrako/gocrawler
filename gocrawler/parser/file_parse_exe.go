package parser

import (
	"github.com/PuerkitoBio/goquery"
	"gocrawler/db"
	"gocrawler/util/strutil"
	"strings"
	"gocrawler/bean"
)

type ExeFileParser struct {
}

func (p *ExeFileParser) StartUrl() string {
	//return "http://www.onlinedown.net/"
	return "http://rj.baidu.com/index.html"
}

func (p *ExeFileParser) FileFilter(url string) bool {
	return true
}

func (p *ExeFileParser) FileParse(doc *goquery.Document, db *db.FileDB) {
	s := doc.Find("head").Find("title").First()
	name := s.Text()

	// 爬所有链接
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		v, _ := s.Attr("href")

		if !strings.HasSuffix(v, ".exe") {
			return
		}

		b := new(bean.FileBean)
		b.Name = name
		b.Url = doc.Url.String()
		b.Suffix = strutil.Suffix(v)
		b.Download = v

		end := strutil.LastIndex(v, ".")
		if end < 0 {
			return
		}
		start := strutil.LastIndex(v, "/")
		if start < 0 {
			return
		}
		//substr := strutil.SubString(v, start+1, end)
		//name, err := url.QueryUnescape(substr)
		//if err != nil {
		//	name = substr // 不需要解码
		//}
		//b.Name = name
		b.Type = "exe"
		db.ReplaceFile(b)
	})
}

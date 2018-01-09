package parser

import (
	"github.com/PuerkitoBio/goquery"
	"gocrawler/db"
)

type FileParser struct {
	BaseParser
	myDB *db.FileDB

	parserImp IFileParser
}

func (p *FileParser) SetupData() {
	p.id = "file"
	p.myDB = db.NewFileDB()
	//p.parserImp = new(Java1234FileParser)
	p.parserImp = new(Java1234FileParser)
	p.startUrl = p.parserImp.StartUrl()
}

func (p *FileParser) Release() {
	if p.myDB != nil {
		p.myDB.Close()
		p.myDB = nil
	}
}

func (p *FileParser) Filter(url string) bool {
	if !p.BaseParser.Filter(url) {
		return false
	}
	if !p.parserImp.FileFilter(url) {
		return false
	}
	return true
}

func (p *FileParser) Parse(doc *goquery.Document) []string {
	urls := p.BaseParser.parseHref(doc)
	p.doParse(doc)
	return urls
}

func (p *FileParser) doParse(doc *goquery.Document) {
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	p.parserImp.FileParse(doc, p.myDB)
}

//func (p *FileParser) parseHref(doc *goquery.Document) {
//	if doc == nil {
//		return
//	}
//
//	// 爬所有链接
//	doc.Find("a").Each(func(i int, s *goquery.Selection) {
//		v, _ := s.Attr("href")
//
//		if p.pdfFilter(v) {
//			b := new(bean.FileBean)
//			b.Url = doc.Url.Host + doc.Url.String()
//			b.Suffix = strutil.Suffix(v)
//			b.Download = v
//
//			end := strutil.LastIndex(v, ".")
//			if end < 0 {
//				return
//			}
//			start := strutil.LastIndex(v, "/")
//			if start < 0 {
//				return
//			}
//			substr := strutil.SubString(v, start+1, end)
//			name, err := url.QueryUnescape(substr)
//			if err != nil {
//				name = substr // 不需要解码
//			}
//			b.Name = name
//			p.myDB.ReplaceFile(b)
//		}
//	})
//}
//
//func (p *FileParser) pdfFilter(url string) bool {
//	if strings.HasSuffix(url, ".pdf") {
//		return true
//	}
//	return false
//}
//
//func (p *FileParser) apeFilter(url string) bool {
//	if strings.Contains(url, "www.baidu.com") { // 盗链
//		return false
//	}
//
//	if strings.HasSuffix(url, ".ape") || strings.HasSuffix(url, ".flac") || strings.HasSuffix(url, ".wav") {
//		return true
//	}
//	return false
//}

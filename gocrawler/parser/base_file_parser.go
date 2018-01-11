package parser

import (
	"gocrawler/db"
)

type BaseFileParser struct {
	BaseParser
	myDB *db.FileDB

	//parserImp IFileParser
}

func (p *BaseFileParser) SetupData() {
	p.id = "file"
	p.myDB = db.NewFileDB()
}

func (p *BaseFileParser) Release() {
	if p.myDB != nil {
		p.myDB.Close()
		p.myDB = nil
	}
}

func (p *BaseFileParser) Filter(url string) bool {
	if !p.BaseParser.Filter(url) {
		return false
	}
	return true
}

//func (p *BaseFileParser) Parse(doc *goquery.Document) []string {
//	urls := p.BaseParser.parseHref(doc)
//	//p.doParse(doc)
//	return urls
//}

//func (p *BaseFileParser) doParse(doc *goquery.Document) {
//	defer func() {
//		if err := recover(); err != nil {
//		}
//	}()
//}

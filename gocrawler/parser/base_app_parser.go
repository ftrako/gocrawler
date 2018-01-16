package parser

import (
	"gocrawler/bean"
	"gocrawler/db"
)

type BaseAppParser struct {
	BaseParser

	os         string // android or ios
	storeId    string
	storeName  string
	categories map[string]*bean.CategoryBean
	myDB       *db.AppDB
}

func (p *BaseAppParser) SetupData() {
	p.myDB = db.NewAppDB()
	p.categories = make(map[string]*bean.CategoryBean)
}

func (p *BaseAppParser) Release() {
	if p.myDB != nil {
		p.myDB.Close()
		p.myDB = nil
	}
}

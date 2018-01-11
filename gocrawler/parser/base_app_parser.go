package parser

import "gocrawler/db"

type BaseAppParser struct {
	BaseParser

	os        string // android or ios
	storeId   string
	storeName string
	myDB      *db.AppDB
}

func (p *BaseAppParser) SetupData() {
	p.myDB = db.NewAppDB()
}

func (p *BaseAppParser) Release() {
	if p.myDB != nil {
		p.myDB.Close()
		p.myDB = nil
	}
}

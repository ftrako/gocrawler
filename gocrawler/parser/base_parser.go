package parser

import "gocrawler/data"

type BaseParser struct {
	Os        string // android or ios
	StoreId   string
	StoreName string
	StartUrl  string
	UrlQueue  *data.UrlQueue
}

func (p *BaseParser) GetStartUrl() string {
	return p.StartUrl
}

func (p *BaseParser) GetStoreId() string {
	return p.StoreId
}

func (p *BaseParser) GetOs() string {
	return p.Os
}

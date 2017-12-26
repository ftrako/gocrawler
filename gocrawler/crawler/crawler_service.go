package crawler

import (
	"sync"
)

type CrawlerService struct {
	crawlers map[string]*Crawler
}

var crawlerServiceInstance *CrawlerService
var crawlerServiceLock = &sync.Mutex {}

func SharedService() *CrawlerService {
	if crawlerServiceInstance == nil {
		crawlerServiceLock.Lock()
		defer crawlerServiceLock.Unlock()
		crawlerServiceInstance = newCrawlerService()
	}

	return crawlerServiceInstance
}

func (p *CrawlerService) StartOneCrawler(parserId string) {
	if _,ok := p.crawlers[parserId];!ok {
		p.crawlers[parserId] = NewCrawler(parserId)
	}
	p.crawlers[parserId].Start()
}

func newCrawlerService() *CrawlerService {
	p := new(CrawlerService)
	p.crawlers = make(map[string]*Crawler)
	return p
}
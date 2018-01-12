package crawler

import (
	"gocrawler/parser"
	"runtime"
	"sync"
	"time"
)

type CrawlerService struct {
	crawlers map[string]*Crawler
	max      int // 最大crawler数量限制，暂时只支持单个爬虫对象

	timerForGC *time.Ticker
}

var crawlerServiceInstance *CrawlerService
var crawlerServiceLock = &sync.Mutex{}

func SharedService() *CrawlerService {
	if crawlerServiceInstance == nil {
		crawlerServiceLock.Lock()
		defer crawlerServiceLock.Unlock()
		crawlerServiceInstance = new(CrawlerService)
		crawlerServiceInstance.setupData()
	}

	return crawlerServiceInstance
}

func (p *CrawlerService) RestartOneCrawler(parserType parser.ParserType) bool {
	return p.startOneCrawler(parserType, true)
}

func (p *CrawlerService) StartOneCrawler(parserType parser.ParserType) bool {
	return p.startOneCrawler(parserType, false)
}

// 目前只支持一个爬虫器
func (p *CrawlerService) startOneCrawler(parserType parser.ParserType, restart bool) bool {
	if len(p.crawlers) >= p.max { // 超过最大限制
		return false
	}

	strType := string(parserType)
	if _, ok := p.crawlers[strType]; ok { // 已经启动
		return false
	}
	c := NewCrawler(parserType, !restart)
	if c == nil {
		return false
	}
	p.crawlers[strType] = c
	p.crawlers[strType].Start()
	p.crawlers[strType].Release()
	delete(p.crawlers, strType)
	return true
}

func (p *CrawlerService) Release() {
	if p.timerForGC != nil {
		p.timerForGC.Stop()
	}
	for k, v := range p.crawlers {
		if v != nil {
			v.Release()
		}
		delete(p.crawlers, k)
	}
}

func (p *CrawlerService) setupData() {
	p.crawlers = make(map[string]*Crawler)
	p.max = 1
	go p.toggleFreeMemory()
}

func (p *CrawlerService) toggleFreeMemory() {
	ticker := time.NewTicker(10 * time.Minute)
	for range ticker.C {
		go runtime.GC()
	}
}

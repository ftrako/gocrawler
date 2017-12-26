package crawler

import (
	"fmt"
	"gocrawler/data"
	"gocrawler/parser"
	"sync"
	"github.com/PuerkitoBio/goquery"
	"time"
)

type Crawler struct {
	urlQueue  *data.UrlQueue
	appParser parser.IParser
	count     uint64
	waitGroup sync.WaitGroup
	locker    sync.Mutex
}

func NewCrawler(parserId string) *Crawler {
	p := new(Crawler)
	p.urlQueue = data.NewUrlQueue()
	p.appParser = parser.NewParser(p.urlQueue, parserId)
	p.count = 0
	return p
}

func (p *Crawler) Start() {
	p.urlQueue.AddNewUrl(p.appParser.GetStartUrl())
	p.toggleWorks()
	p.waitGroup.Wait()
}

func (p *Crawler) toggleWorks() {
	p.locker.Lock()
	defer p.locker.Unlock()
	minCount := p.urlQueue.GetUsableCount()
	for loop := 0; loop < minCount; loop++ {
		go p.toggleOneWork()
		p.waitGroup.Add(1)
	}
}

func (p *Crawler) toggleOneWork() {
	defer func() {
		if !p.urlQueue.Empty() {
			p.toggleWorks()
		}
		p.waitGroup.Done()
	}()

	url := p.urlQueue.ToggleRunUrl()

	p.doWork(url)

	p.count++
	p.urlQueue.DoneUrl(url)
	fmt.Println("toggle work", p.count, time.Now(), url)
}

func (p *Crawler) doWork(url string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:", err) // 这里的err其实就是panic传入的内容
		}
	}()
	if url == "" {
		return
	}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return
	}

	p.appParser.Parse(doc)
}

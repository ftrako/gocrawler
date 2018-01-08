package crawler

import (
	"gocrawler/parser"
	"sync"
	"github.com/PuerkitoBio/goquery"
	"runtime"
	"fmt"
	"time"
	"strings"
	"gocrawler/urlmgr"
	"gocrawler/log"
	"gocrawler/conf"
)

type Crawler struct {
	urlQueue *urlmgr.UrlQueue
	parser   parser.IParser
	//count     uint64 // 抓的条数
	waitGroup sync.WaitGroup
	//locker    sync.Mutex

	log *log.FileLog
}

func NewCrawler(parserType parser.ParserType, resume bool) *Crawler {
	p := new(Crawler)
	p.SetupData(parserType, resume)
	return p
}

func (p *Crawler) SetupData(parserType parser.ParserType, resume bool) {
	p.parser = parser.NewParser(parserType)
	p.urlQueue = urlmgr.NewUrlQueue(p.parser.GetId(), resume)

	year, month, day := time.Now().Date()
	logName := fmt.Sprintf("trace_%04d%02d%02d.log", year, month, day)
	p.log = log.NewFileLog(conf.GetDataPath() + "/" + logName)
}

func (p *Crawler) Start() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	p.waitGroup.Add(1)
	p.urlQueue.AddNewUrl(p.parser.GetStartUrl())
	p.toggleWork()
	p.waitGroup.Wait()

	p.urlQueue.Release()
}

func (p *Crawler) toggleWork() {
	defer func() {
		if p.urlQueue.Empty() {
			p.waitGroup.Done()
		}
	}()
	for {
		url := p.urlQueue.GetWaitUrl()
		if url == "" {
			break
		}

		go p.doWork(url)
	}
}

func (p *Crawler) doWork(url string) {
	defer func() {
		if err := recover(); err != nil {
		}
		p.urlQueue.DoneUrl(url)
		p.toggleWork()
	}()
	if url == "" {
		return
	}

	startTime := time.Now().Unix()

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return
	}

	// 拉取网页时间过长，记录下来
	endTime := time.Now().Unix()
	if endTime-startTime > 5 {
		p.log.Println(fmt.Sprintf("long time %0d , url = %s", (endTime - startTime), url))
	}

	urls := p.parser.Parse(doc)

	for _, v := range urls {
		if v == "" {
			continue
		}
		if !strings.HasPrefix(v, "http://") && !strings.HasPrefix(v, "https://") { // 支持内部链接
			if strings.HasPrefix(v, "/") {
				v = p.parser.GetHost() + v
			} else {
				v = p.parser.GetHost() + "/" + v
			}
		}
		if p.urlQueue.Exist(v) { // 先检查是否已存在，再执行过滤器，会提高性能
			continue
		}
		if p.parser.Filter(v) { // 过滤不满足条件的url
			p.urlQueue.AddNewUrl(v)
		}
	}
}

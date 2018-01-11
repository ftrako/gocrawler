package crawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gocrawler/conf"
	"gocrawler/log"
	"gocrawler/parser"
	"gocrawler/urlmgr"
	"runtime"
	"strings"
	"sync"
	"time"
)

type Crawler struct {
	urlQueue  *urlmgr.UrlQueue
	parser    parser.IParser
	waitGroup sync.WaitGroup

	log *log.FileLog
}

func NewCrawler(parserType parser.ParserType, resume bool) *Crawler {
	p := new(Crawler)
	res := p.SetupData(parserType, resume)
	if !res {
		return nil
	}
	return p
}

func (p *Crawler) SetupData(parserType parser.ParserType, resume bool) bool {
	p.parser = parser.NewParser(parserType)
	if p.parser == nil {
		return false
	}
	p.urlQueue = urlmgr.NewUrlQueue(p.parser.GetId(), resume)

	year, month, day := time.Now().Date()
	logName := fmt.Sprintf("trace_%04d%02d%02d.log", year, month, day)
	p.log = log.NewFileLog(conf.GetDataPath() + "/log/" + logName)
	return true
}

func (p *Crawler) Release() {
	if p.parser != nil {
		p.parser.Release()
		p.parser = nil
	}
	if p.urlQueue != nil {
		p.urlQueue.Release()
		p.urlQueue = nil
	}
}

func (p *Crawler) Start() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	p.waitGroup.Add(1)
	p.urlQueue.AddNewUrl(p.parser.GetStartUrl())
	p.toggleWork()
	p.waitGroup.Wait()
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
			fmt.Println(err, url)
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
		//p.log.Println("error: failed new doc " + err.Error())
		return
	}
	endTime := time.Now().Unix()
	if endTime-startTime > 5 {
		p.log.Println(fmt.Sprintf("long time %0d , url = %s", endTime-startTime, url))
	}

	//doc, err := goquery.NewDocument(url)
	//timeout := time.Second * 20
	//resp, err := httputil.DoGetWithTimeout(url, timeout)
	//if err != nil {
	//	fmt.Println("error:", url, time.Now())
	//	p.log.Println("timeout " + timeout.String() + ", url " + url)
	//	return
	//}

	//doc, err2 := goquery.NewDocumentFromResponse(resp)
	//if err2 != nil {
	//	fmt.Println("error:", err2)
	//	return
	//}

	urls := p.parser.Parse(doc)
	for _, v := range urls {
		if v == "" {
			continue
		}
		if !strings.HasPrefix(v, "http://") && !strings.HasPrefix(v, "https://") { // 支持内部链接
			preUrl := doc.Url.Scheme + "://" + doc.Url.Host // 自动添加host组成一个完整的url
			if doc.Url.Port() != "" {
				preUrl += ":" + doc.Url.Port()
			}

			if !strings.HasPrefix(v, "/") {
				preUrl += "/"
			}

			v = preUrl + v
		}
		if p.urlQueue.Exist(v) { // 先检查是否已存在，再执行过滤器，会提高性能
			continue
		}
		if p.parser.Filter(v) { // 过滤不满足条件的url
			p.urlQueue.AddNewUrl(v)
		}
	}
}

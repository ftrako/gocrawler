package urlmgr

import (
	"container/list"
	"fmt"
	"gocrawler/backup"
	"gocrawler/conf"
	"gocrawler/util/cryptutil"
	"gocrawler/util/timeutil"
	"sync"
	"time"
)

type UrlQueue struct {
	waitList *list.List        // FIFO
	waitMap  map[string]string // key为md5(url),value为url
	doneMap  map[string]string // key为md5(url),value为url

	name string

	//waitLocker  sync.Mutex
	//doneLocker  sync.Mutex
	//countLocker sync.Mutex

	locker sync.Mutex // 所有操作用一个锁

	runCountMax int // 最多运行线程数
	runCount    int // 当前运行线程数量

	waitBak *backup.Backup
	doneBak *backup.Backup

	count          int
	timerForBackup *time.Ticker
}

func NewUrlQueue(name string, resume bool) *UrlQueue {
	queue := new(UrlQueue)
	queue.SetupData(name, resume)
	return queue
}

func (p *UrlQueue) SetupData(name string, resume bool) {
	p.waitMap = make(map[string]string)
	p.doneMap = make(map[string]string)
	p.waitList = list.New()
	p.runCountMax = 10
	p.name = name
	p.waitBak = backup.NewBackup(conf.GetDataPath() + "/url/waiturlmap_" + p.name + ".dat")
	p.doneBak = backup.NewBackup(conf.GetDataPath() + "/url/doneurlmap_" + p.name + ".dat")

	if resume {
		p.loadBackup()
	}
	go p.toggleBackup()
}

func (p *UrlQueue) SetExecCountMax(max int) {
	if max <= 0 {
		return
	}
	p.runCountMax = max
}

// AddNewUrl none->wait状态切换
func (p *UrlQueue) AddNewUrl(url string) {
	if url == "" || p.Exist(url) {
		return
	}

	//p.waitLocker.Lock()
	//defer p.waitLocker.Unlock()
	p.locker.Lock()
	defer p.locker.Unlock()

	md5 := cryptutil.MD5(url)
	p.waitList.PushBack(url)
	p.waitMap[md5] = url

	//fmt.Println("add new wait len",len(p.waitMap),", done len",len(p.doneMap), ", new url", url)
}

// 获取等待队列中的url，自动同步从等待队列中移除
func (p *UrlQueue) GetWaitUrl() string {
	//p.waitLocker.Lock()
	//defer p.waitLocker.Unlock()
	//p.countLocker.Lock()
	//defer p.countLocker.Unlock()
	//p.doneLocker.Lock()
	//defer p.doneLocker.Unlock()
	p.locker.Lock()
	defer p.locker.Unlock()

	if len(p.waitMap) <= 0 { // 等待队列没有数据了
		return ""
	}
	if p.runCount > p.runCountMax { // 运行池已满
		return ""
	}

	for {
		ele := p.waitList.Front()
		if ele == nil {
			return ""
		}
		url := ele.Value.(string)
		md5 := cryptutil.MD5(url)

		if p.doneMap[md5] != "" {
			continue
		}
		p.runCount++

		p.waitList.Remove(ele)
		//delete(p.waitMap, md5)
		return url
	}

	return ""
}

// DoneUrl run->done状态切换
func (p *UrlQueue) DoneUrl(url string) {
	//p.doneLocker.Lock()
	//defer p.doneLocker.Unlock()
	//p.countLocker.Lock()
	//defer p.countLocker.Unlock()
	//p.waitLocker.Lock()
	//p.waitLocker.Unlock()
	p.locker.Lock()
	defer p.locker.Unlock()

	md5 := cryptutil.MD5(url)
	delete(p.waitMap, md5)
	p.doneMap[md5] = url
	p.runCount--

	fmt.Println("doneurl wait len", len(p.waitMap), ",run len", p.runCount, ", done len", len(p.doneMap), ",", timeutil.TimeStr(time.Now()), ", done url", url)
}

// Exist true表示队列中已存在
func (p *UrlQueue) Exist(url string) bool {
	//p.waitLocker.Lock()
	//defer p.waitLocker.Unlock()
	//p.doneLocker.Lock()
	//defer p.doneLocker.Unlock()

	p.locker.Lock()
	defer p.locker.Unlock()

	md5 := cryptutil.MD5(url)
	if _, ok := p.waitMap[md5]; ok {
		return true
	}
	if _, ok := p.doneMap[md5]; ok {
		return true
	}
	return false
}

func (p *UrlQueue) Empty() bool {
	//p.waitLocker.Lock()
	//defer p.waitLocker.Unlock()
	//p.countLocker.Lock()
	//defer p.countLocker.Unlock()
	p.locker.Lock()
	defer p.locker.Unlock()

	return (len(p.waitMap) + p.runCount) <= 0
}

func (p *UrlQueue) Release() {
	if p.timerForBackup != nil {
		p.timerForBackup.Stop()
	}
	p.backup()
}

func (p *UrlQueue) toggleBackup() {
	p.timerForBackup = time.NewTicker(20 * time.Minute)
	for range p.timerForBackup.C {
		go p.backup()
	}
}

func (p *UrlQueue) loadBackup() {
	//p.waitLocker.Lock()
	//defer p.waitLocker.Unlock()
	//p.doneLocker.Lock()
	//defer p.doneLocker.Unlock()

	p.locker.Lock()
	defer p.locker.Unlock()

	d := p.waitBak.LoadData()
	if d != nil {
		for k, v := range d {
			p.waitMap[k] = v
			p.waitList.PushBack(v)
		}
	}

	//p.doneLocker.Lock()
	//defer p.doneLocker.Unlock()
	d = p.doneBak.LoadData()
	if d != nil {
		for k, v := range d {
			p.doneMap[k] = v
		}
	}

	fmt.Println("load backup ok wait =", len(p.waitMap), ", done =", len(p.doneMap))
}

func (p *UrlQueue) backup() {
	//p.waitLocker.Lock()
	//defer p.waitLocker.Unlock()
	//p.doneLocker.Lock()
	//defer p.doneLocker.Unlock()
	//p.countLocker.Lock()
	//defer p.countLocker.Unlock()
	p.locker.Lock()
	defer p.locker.Unlock()

	p.waitBak.Backup(p.waitMap)
	p.doneBak.Backup(p.doneMap)
	p.count++

	fmt.Println("backup wait len", len(p.waitMap), ", done len", len(p.doneMap), ", count", p.count)
}

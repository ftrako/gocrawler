package data

import (
	"container/list"
	"gocrawler/util/cryptutil"
	"sync"
	"fmt"
)

type UrlQueue struct {
	waitQueue *list.List //
	// runQueue  list.List      //
	waitMap map[string]bool // key为url
	runMap  map[string]bool // key为url
	doneMap map[string]bool // key为md5(url)

	waitQueueLock sync.Mutex
	runQueueLock  sync.Mutex
	doneQueueLock sync.Mutex
	locker sync.Mutex

	runCountMax int // 最多运行线程数
}

func NewUrlQueue() *UrlQueue {
	queue := new(UrlQueue)
	queue.waitMap = make(map[string]bool)
	queue.runMap = make(map[string]bool)
	queue.doneMap = make(map[string]bool)
	queue.waitQueue = list.New()
	queue.runCountMax = 10
	return queue
}

func (p *UrlQueue) SetExecCountMax(max int) {
	if max <= 0 {
		return
	}
	p.runCountMax = max
}

//func (p *UrlQueue) GetWorkableCount() int {
//	p.runQueueLock.Lock()
//	defer p.runQueueLock.Unlock()
//	p.locker.Lock()
//	defer p.locker.Unlock()
//	return p.runCountMax - len(p.runMap)
//}

func (p *UrlQueue) GetUsableCount() int {
	p.locker.Lock()
	defer p.locker.Unlock()
	waitCount := len(p.waitMap)
	runCount := p.runCountMax - len(p.runMap)
	minCount := runCount
	if minCount > waitCount {
		minCount = waitCount
	}
	fmt.Println("mincount",minCount,",workablecount",runCount,", waitcount",waitCount)
	return minCount
}

//func (p *UrlQueue) GetWaitCount() int {
//	p.waitQueueLock.Lock()
//	defer p.waitQueueLock.Unlock()
//	return len(p.waitMap)
//}

// AddNewUrl none->wait状态切换
func (p *UrlQueue) AddNewUrl(url string) {
	if url == "" || p.Exist(url) {
		return
	}

	p.locker.Lock()
	defer p.locker.Unlock()

	//p.waitQueueLock.Lock()
	//defer p.waitQueueLock.Unlock()
	p.waitQueue.PushBack(url)
	p.waitMap[url] = false
}

// ToggleRunUrl wait->run状态切换
func (p *UrlQueue) ToggleRunUrl() string {
	p.locker.Lock()
	defer p.locker.Unlock()
	//p.waitQueueLock.Lock()
	//p.runQueueLock.Lock()
	//defer p.waitQueueLock.Unlock()
	//defer p.runQueueLock.Unlock()

	if len(p.waitMap) <= 0 { // 等待队列没有数据了
		return ""
	}
	if len(p.runMap) >= p.runCountMax { // 运行池已满
		return ""
	}

	ele := p.waitQueue.Front()
	url := ele.Value.(string)
	p.waitQueue.Remove(ele)
	delete(p.waitMap, url)

	p.runMap[url] = false

	return url
}

// DoneUrl run->done状态切换
func (p *UrlQueue) DoneUrl(url string) {
	p.locker.Lock()
	defer p.locker.Unlock()
	//p.runQueueLock.Lock()
	//p.doneQueueLock.Lock()
	//defer p.runQueueLock.Unlock()
	//defer p.doneQueueLock.Unlock()

	delete(p.runMap, url)
	fmt.Println("DoneUrl, workablecount",len(p.runMap))

	md5 := cryptutil.MD5(url)
	p.doneMap[md5] = false
}

// Exist true表示队列中已存在
func (p *UrlQueue) Exist(url string) bool {

	p.locker.Lock()
	defer p.locker.Unlock()
	//p.waitQueueLock.Lock()
	//p.runQueueLock.Lock()
	//p.doneQueueLock.Lock()
	//defer p.waitQueueLock.Unlock()
	//defer p.runQueueLock.Unlock()
	//defer p.doneQueueLock.Unlock()

	if _, ok := p.waitMap[url]; ok {
		return true
	}

	if _, ok := p.runMap[url]; ok {
		return true
	}

	md5 := cryptutil.MD5(url)
	if _, ok := p.doneMap[md5]; ok {
		return true
	}
	return false
}

func (p *UrlQueue) Empty() bool {

	p.locker.Lock()
	defer p.locker.Unlock()
	//p.waitQueueLock.Lock()
	//defer p.waitQueueLock.Unlock()
	return len(p.waitMap) <= 0
}
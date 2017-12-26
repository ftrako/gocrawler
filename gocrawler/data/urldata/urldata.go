package urldata

// import (
// 	"container/list"
// 	"sync"
// )

// // 抓取的网址，bool true表示已抓取过
// // var waitUrls = make(map[string]bool)
// // var workingUrls = make(map[string]bool)

// var waitQueue = list.New()
// var runQueue = list.New()

// var worksMap = make(map[string]int)

// var workMaxCount = 10 // 最多10个work

// var queueLock sync.Mutex

// // func Exist(url string) bool {
// // 	return false
// // }

// // func Add(url string) {

// // }

// // AddURL 添加url
// func AddURL(url string) {
// 	defer queueLock.Unlock()
// 	queueLock.Lock()

// 	if _, ok := worksMap[url]; ok {
// 		return
// 	}
// 	worksMap[url] = 0
// 	waitQueue.PushBack(url)
// }

// // DoOneWork 等待队列移到运行队列
// func DoOneWork() *list.Element {
// 	defer queueLock.Unlock()

// 	queueLock.Lock()
// 	if runQueue.Len() >= workMaxCount || waitQueue.Len() <= 0 {
// 		return nil
// 	}

// 	ele := waitQueue.Front()
// 	waitQueue.Remove(ele)
// 	runQueue.MoveToBack(ele)
// 	return ele
// }

// // DoneOneWork
// func DoneOneWork(ele *list.Element) {
// 	defer queueLock.Unlock()
// 	queueLock.Lock()
// 	runQueue.Remove(ele)
// }

// // Exist url是否已在队列中
// func Exist(url string) bool {
// 	if _, ok := worksMap[url]; ok {
// 		return true
// 	}
// 	return false
// }

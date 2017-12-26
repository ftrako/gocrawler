package crawler
//
//import (
//	"fmt"
//	"gocrawler/data"
//	"gocrawler/db"
//	"gocrawler/parser"
//	"sync"
//	"time"
//
//	"github.com/PuerkitoBio/goquery"
//)
//
//var queue sync.WaitGroup
//
//var count uint64 = 0
//
//var urldata = data.NewUrlQueue()
//var appParser = parser.NewParser(urldata, "appstore")
//
//// 爬网页
//
//// type BaseCrawler interface {
//// 	StartCartch(url string)
//// }
//
//// StartCrawlerService 启动抓取服务
//func StartCrawlerService() {
//	db.Open()
//	startNewCrawler()
//	db.Close()
//}
//
//func startNewCrawler() {
//	addURL(appParser.GetStartUrl())
//	toggleOneWork()
//	queue.Wait()
//}
//
//func addURL(url string) {
//	urldata.AddNewUrl(url)
//}
//
//func toggleOneWork() {
//	url := urldata.ToggleRunUrl()
//	if url == "" {
//		return
//	}
//
//	count++
//	fmt.Println("toggle work", count, time.Now())
//
//	go doWork(&queue, url)
//	queue.Add(1)
//}
//
//func doWork(queue *sync.WaitGroup, url string) {
//	defer func() {
//		if err := recover(); err != nil {
//			fmt.Println("err:", err) // 这里的err其实就是panic传入的内容
//		}
//		urldata.DoneUrl(url)
//		toggleOneWork()
//		queue.Done()
//	}()
//	if url == "" {
//		return
//	}
//	doc, err := goquery.NewDocument(url)
//	if err != nil {
//		return
//	}
//
//	appParser.Parse(doc)
//	// parseDoc(doc)
//}
//
//// // 解析网页
//// func parseDoc(doc *goquery.Document) {
//// 	if doc == nil {
//// 		return
//// 	}
//
//// 	// 爬所有链接
//// 	doc.Find("a").Each(func(i int, s *goquery.Selection) {
//// 		// v, _ := s.Attr("href")
//// 		// addURL(v)
//// 	})
//
//// 	// 爬分类
//// 	parseCategory(doc)
//
//// 	// 爬应用
//// 	parseApp(doc)
//// }
//
//// func parseCategory(doc *goquery.Document) {
//// 	if doc == nil {
//// 		return
//// 	}
//
//// 	doc.Find("li.app-tag-wrap").Find("a.app-tag").Find("span").Each(func(i int, s *goquery.Selection) {
//// 		// fmt.Println(s.Text())
//// 		// 插入第一层
//// 		var b bean.CategoryBean
//// 		b.Name = s.Text()
//// 		db.ReplaceCategory(&b)
//// 		parseSubCategory(doc, "li.app-tag-wrap", s.Text())
//// 	})
//
//// 	doc.Find("li.game-tag-wrap").Find("a.game-tag").Find("span").Each(func(i int, s *goquery.Selection) {
//// 		// fmt.Println(s.Text())
//// 		// 插入第一层
//// 		var b bean.CategoryBean
//// 		b.Name = s.Text()
//// 		db.ReplaceCategory(&b)
//// 		parseSubCategory(doc, "li.game-tag-wrap", s.Text())
//// 	})
//// }
//
//// func parseSubCategory(doc *goquery.Document, basequery string, basename string) {
//// 	if doc == nil {
//// 		return
//// 	}
//// 	doc.Find(basequery).Find("li.parent-cate").Each(func(i int, subs *goquery.Selection) {
//// 		subs.Find("a.cate-link").Each(func(j int, subss *goquery.Selection) {
//// 			// fmt.Println(basename + "-->" + subss.Text())
//// 			subs.Find("li.child-cate").Each(func(j int, subsss *goquery.Selection) {
//// 				subsss.Find("a").Each(func(k int, subssss *goquery.Selection) {
//// 					// fmt.Println(basename + "-->" + subss.Text() + "-->" + subssss.Text())
//// 					var b bean.CategoryBean
//// 					b.SuperName = subss.Text()
//// 					b.StoreId = "wandoujia"
//// 					b.Name = subssss.Text()
//// 					db.ReplaceCategory(&b) // 插入分类
//// 				})
//// 			})
//// 		})
//// 	})
//// }
//
//// func parseApp(doc *goquery.Document) {
//// 	var b bean.AppBean
//// 	b.Os = "android"
//// 	doc.Find("div.detail-wrap").Find("div.detail-top.clearfix").Find("div.app-info").Find("a[data-app-id]").Each(func(j int, s *goquery.Selection) {
//// 		var text string
//// 		text, _ = s.Attr("data-app-name")
//// 		if text == "" {
//// 			return // 没有名称
//// 		}
//// 		b.Name = text
//// 		text, _ = s.Attr("data-app-pname")
//// 		if text == "" {
//// 			return // 没有appid
//// 		}
//// 		b.AppId = text
//// 		text, _ = s.Attr("data-app-vname")
//// 		b.Version = text
//// 		text, _ = s.Attr("data-install")
//// 		b.InstallCount = text
//// 	})
//// 	doc.Find("div.detail-wrap").Find("div.infos").Find("dl.infos-list").Each(func(j int, s *goquery.Selection) {
//// 		s.Find("meta[itemprop][content]").Each(func(j int, ss *goquery.Selection) {
//// 			text, _ := ss.Attr("content")
//// 			b.Size = text
//// 		})
//// 		var categories string
//// 		s.Find("dd.tag-box").Find("a").Each(func(j int, ss *goquery.Selection) {
//// 			categories += ss.Text() + ";"
//// 		})
//// 		b.Category = categories
//// 		s.Find("time[datetime]").Each(func(j int, ss *goquery.Selection) {
//// 			b.UpdateTime = ss.Text()
//// 		})
//// 		s.Find("span.dev-sites").Each(func(j int, ss *goquery.Selection) {
//// 			b.Vender = ss.Text()
//// 		})
//// 		s.Find("dd.perms[itemprop=operatingSystems]").Each(func(j int, ss *goquery.Selection) {
//// 			text := ss.Text()
//// 			text = strings.Trim(text, "\n")
//// 			index := strings.Index(text, "\n")
//// 			text = stringutil.SubString(text, 0, index)
//// 			text = strings.TrimSpace(text)
//// 			b.MinVersion = text
//// 		})
//// 	})
//
//// 	db.ReplaceApp(&b)
//// }

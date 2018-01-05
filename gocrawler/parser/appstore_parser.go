package parser

import (
	"encoding/json"
	"gocrawler/bean"
	"gocrawler/db"
	"gocrawler/util/strutil"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type AppStoreParser struct {
	BaseParser

	os        string // android or ios
	storeId   string
	storeName string
	myDB *db.AppDB
}

func (p *AppStoreParser) Filter(url string) bool {
	if !strings.Contains(url, "apple.com") {
		return false
	}
	if !p.BaseParser.Filter(url) {
		return false
	}
	return true
}

func (p *AppStoreParser) Parse(doc *goquery.Document) []string{
	var urls = make([]string, 0, 100)
	if doc == nil {
		return urls
	}

	// 爬所有链接
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		v, _ := s.Attr("href")
		urls = append(urls, v)
	})

	// 爬分类
	p.parseCategory(doc)

	// 爬应用
	p.parseApp(doc)

	return urls
}

func (p *AppStoreParser) parseCategory(doc *goquery.Document) {
	if doc == nil {
		return
	}

	var level0Name string

	// 总共三级

	// 第一级
	doc.Find("div#content").Find("div#media-type-nav.nav").Find("ul.list").Each(func(i int, s *goquery.Selection) {
		s.Find("li").Find("a").Each(func(i int, s2 *goquery.Selection) {
			var c bean.CategoryBean
			c.Name = s2.Text()
			c.StoreId = p.storeId
			c.StoreName = p.storeName
			p.myDB.ReplaceCategory(&c)
		})

		// 当前selected项
		s2 := s.Find("li.selected").First()
		level0Name = s2.Text()
	})

	// 第二级
	doc.Find("div#content").Find("div.main.nav#genre-nav").Find("div").Find("ul.list.column").Find("li").Each(func(i int, s *goquery.Selection) {
		var level1Name string
		s2 := s.Find("a.top-level-genre").First()
		level1Name = s2.Text()
		var c bean.CategoryBean
		c.Name = s2.Text()
		c.StoreId = p.storeId
		c.StoreName = p.storeName
		c.SuperName = level0Name
		p.myDB.ReplaceCategory(&c)

		// 第三级
		s.Find("ul").Find("li").Find("a").Each(func(i int, s2 *goquery.Selection) {
			var c2 bean.CategoryBean
			c2.Name = s2.Text()
			c2.StoreId = p.storeId
			c2.StoreName = p.storeName
			c2.SuperName = level1Name
			p.myDB.ReplaceCategory(&c2)
		})

	})
}

func (p *AppStoreParser) parseApp(doc *goquery.Document) {
	s := doc.Find("div#left-stack").Find("a").First()
	text, _ := s.Attr("href")
	if text == "" {
		return
	}

	start := strings.LastIndex(text, "/") + 3
	end := strings.LastIndex(text, "?")
	text = strutil.SubString(text, start, end)

	// lookup应用信息，全球收不到时再搜索中国区
	jsontext := p.requestJsonFile("https://itunes.apple.com/lookup?id=" + text)
	b := p.parseJson(jsontext)
	if b == nil {
		jsontext = p.requestJsonFile("https://itunes.apple.com/cn/lookup?id=" + text)
		b = p.parseJson(jsontext)
	}
	if b == nil {
		return
	}

	b.Os = p.os
	b.StoreId = p.storeId

	// b.Name = doc.Find("div#title.intro").Find("h1[itemprop=name]").First().Text()

	b.Category = ";" + doc.Find("div#left-stack").Find("ul.list").Find("li.genre").Find("span[itemprop=applicationCategory]").First().Text() + ";"

	p.myDB.ReplaceApp(b)
}

func (p *AppStoreParser) requestJsonFile(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)
}

func (p *AppStoreParser) parseJson(jsontext string) *bean.AppBean {
	if jsontext == "" {
		return nil
	}

	var dat map[string]interface{}
	err := json.Unmarshal([]byte(jsontext), &dat)
	if err != nil {
		return nil
	}

	results := dat["results"].([]interface{})
	for _, v := range results {
		newV := v.(map[string]interface{})
		text := newV["bundleId"].(string)
		if text == "" {
			continue
		}
		var b bean.AppBean
		b.AppId = text
		b.IosAppId = strconv.Itoa((int)(newV["trackId"].(float64)))
		b.Name = newV["trackName"].(string)
		b.Vender = newV["artistName"].(string)
		b.MinVersion = newV["minimumOsVersion"].(string)
		b.Version = newV["version"].(string)
		text = newV["fileSizeBytes"].(string)
		size, _ := strconv.Atoi(text)
		b.Size = strconv.Itoa(size/1024.0/1024) + " MB"
		text = newV["currentVersionReleaseDate"].(string)
		index := strings.Index(text, "T")
		text = strutil.SubString(text, 0, index)
		strs := strings.Split(text, "-")
		b.UpdateTime = strs[0] + "年" + strs[1] + "月" + strs[2] + "日"
		return &b
	}
	return nil
}

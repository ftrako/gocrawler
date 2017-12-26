package parser

import (
	"encoding/json"
	"gocrawler/bean"
	"gocrawler/db"
	"gocrawler/util/stringutil"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type AppStoreParser struct {
	BaseParser
}

func (p *AppStoreParser) Parse(doc *goquery.Document) {
	if doc == nil {
		return
	}

	// 爬所有链接
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		v, _ := s.Attr("href")
		p.UrlQueue.AddNewUrl(v)
	})

	// 爬分类
	p.parseCategory(doc)

	// 爬应用
	p.parseApp(doc)
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
			c.StoreId = p.StoreId
			c.StoreName = p.StoreName
			db.ReplaceCategory(&c)
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
		c.StoreId = p.StoreId
		c.StoreName = p.StoreName
		c.SuperName = level0Name
		db.ReplaceCategory(&c)

		// 第三级
		s.Find("ul").Find("li").Find("a").Each(func(i int, s2 *goquery.Selection) {
			var c2 bean.CategoryBean
			c2.Name = s2.Text()
			c2.StoreId = p.StoreId
			c2.StoreName = p.StoreName
			c2.SuperName = level1Name
			db.ReplaceCategory(&c2)
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
	text = stringutil.SubString(text, start, end)

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

	b.Os = p.Os
	b.StoreId = p.StoreId

	// b.Name = doc.Find("div#title.intro").Find("h1[itemprop=name]").First().Text()

	b.Category = ";" + doc.Find("div#left-stack").Find("ul.list").Find("li.genre").Find("span[itemprop=applicationCategory]").First().Text() + ";"

	db.ReplaceApp(b)
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
		text = stringutil.SubString(text, 0, index)
		strs := strings.Split(text, "-")
		b.UpdateTime = strs[0] + "年" + strs[1] + "月" + strs[2] + "日"
		return &b
	}
	return nil
}

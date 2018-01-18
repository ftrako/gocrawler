package parser

import (
	"github.com/PuerkitoBio/goquery"
	"gocrawler/bean"
	"gocrawler/db"
	"gocrawler/util/strutil"
	"strings"
)

type Ape51Parser struct {
	BaseParser

	myDB *db.MusicDB
}

func (p *Ape51Parser) SetupData() {
	p.myDB = db.NewMusicDB()
	p.id = "ape51"
	p.startUrl = "http://www.51ape.com/artist/"
}

func (p *Ape51Parser) Release() {
	if p.myDB != nil {
		p.myDB.Close()
		p.myDB = nil
	}
}

func (p *Ape51Parser) Filter(url string) bool {
	if strings.Contains(url, "51ape.com/zuoci/") ||
		strings.Contains(url, "51ape.com/zuoqu/") ||
		strings.Contains(url, "51ape.com/tags/") ||
		strings.Contains(url, "51ape.com/zhuanji/") ||
		strings.Contains(url, "51ape.com/apeyinyue/") ||
		strings.Contains(url, "51ape.com/flacyinyue/") ||
		strings.Contains(url, "51ape.com/wavyinyue/") ||
		strings.Contains(url, "51ape.com/yueyu/") ||
		strings.Contains(url, "51ape.com/yingwen/") ||
		strings.Contains(url, "51ape.com/经典歌曲/") {
		return false
	}
	if !strings.Contains(url, "51ape.com") {
		return false
	}
	if !p.BaseParser.Filter(url) {
		return false
	}
	return true
}

func (p *Ape51Parser) Parse(doc *goquery.Document) []string {
	urls := p.BaseParser.Parse(doc)
	p.doParse(doc)
	return urls
}

func (p *Ape51Parser) doParse(doc *goquery.Document) {
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	p.parseSong(doc)
}

func (p *Ape51Parser) parseSong(doc *goquery.Document) {
	if doc == nil {
		return
	}

	bean := new(bean.SongBean)

	text := doc.Find("div.fl.over.w638").Find("li.fl.ml_1.mt_08.c999").Text()
	if text == "" {
		return
	}
	bean.Name = text

	text = doc.Find("ul.b_b_s.over").Find("li").Find("a.fl.c3b.ml_1.mt_08").Text()
	if text == "" {
		return
	}
	bean.Singer = text

	index := 0
	doc.Find("h3.c999.fl.mt_05.f_12.n.yh").Each(func(i int, s *goquery.Selection) {
		text = s.Text()
		index++
		if index == 1 {
			subText := "选自专辑《"
			index2 := strutil.Index(text, subText)
			if index2 >= 0 { // 专辑
				text = strutil.SubString(text, index2+strutil.Len(subText), strutil.Len(text)-1)
				bean.Album = text
			}
		} else if index == 3 {
			text = strutil.SubString(text, 1, strutil.Len(text))
			bean.Size = text
		} else if index == 4 {
			text = strutil.SubString(text, 1, strutil.Len(text))
			bean.Language = text
		} else if index == 5 {
			text = strutil.SubString(text, 1, strutil.Len(text))
			bean.Date = text
		}
	})

	text = doc.Find("h1.yh.mt_1.f_32").Text()
	index = strutil.LastIndex(text, ".")
	if index >= 0 {
		bean.Type = strutil.SubString(text, index+1, strutil.Len(text))
	}

	bean.Url = doc.Url.String()

	text, _ = doc.Find("div.fl.over.w638").Find("a").Find("h2.bg_gr.b_b_s.m_s.logo.mt_1.yh.white").Parent().Attr("href")
	bean.Download = text

	text = doc.Find("div.fl.over.w638").Find("b.mt_1.yh.d_b").Text()
	index = strutil.LastIndex(text, "密码：")
	if index >= 0 {
		bean.Code = strutil.SubString(text, index+strutil.Len("密码："), strutil.Len(text))
	}

	p.myDB.ReplaceSong(bean)
}

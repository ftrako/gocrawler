package parser

import (
	"net/http"
	"io/ioutil"
	"gocrawler/util/strutil"
	"strings"
	"fmt"
	"encoding/json"
	"strconv"
	"gocrawler/bean"
)

type AppStoreParseJson struct {
}

func (p *AppStoreParseJson) requestJsonByBundleId(bundleId string) *bean.AppBean {
	// lookup应用信息，全球收不到时再搜索中国区
	urls := []string{"https://itunes.apple.com/lookup?bundleId=" + bundleId, "https://itunes.apple.com/cn/lookup?bundleId=" + bundleId}
	var jsontext string
	for _, v := range urls {
		jsontext = p.requestJson(v)
		if jsontext != "" {
			break
		}
	}
	b := p.parseJson(jsontext)
	return b
}

func (p *AppStoreParseJson) requestJsonByAppId(appId string) *bean.AppBean {
	// lookup应用信息，全球收不到时再搜索中国区
	urls := []string{"https://itunes.apple.com/lookup?id=" + appId, "https://itunes.apple.com/cn/lookup?id=" + appId}
	var jsontext string
	for _, v := range urls {
		jsontext = p.requestJson(v)
		if jsontext != "" {
			break
		}
	}
	b := p.parseJson(jsontext)
	return b
}

func (p *AppStoreParseJson) requestJson(url string) string {
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

func (p *AppStoreParseJson) parseJson(jsontext string) *bean.AppBean {
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

		// 分类
		cids := newV["genreIds"]
		if tmpCids, ok := cids.([]interface{}); ok {
			c := ";"
			for _, v := range tmpCids {
				if v2, ok2 := v.(string); ok2 {
					c += v2 + ";"
				}
			}
			b.Category = c
		}

		fmt.Println("category", b.Category)

		return &b
	}
	return nil
}

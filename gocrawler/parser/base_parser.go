package parser

import (
	"github.com/PuerkitoBio/goquery"
	"gocrawler/util/httputil"
	"gocrawler/util/strutil"
	"strings"
	"time"
)

type BaseParser struct {
	startUrl string
	id       string
}

func (p *BaseParser) GetStartUrl() string {
	return p.startUrl
}

func (p *BaseParser) GetId() string {
	return p.id
}

func (p *BaseParser) Release() {
}

// true表示满足爬虫过滤条件，允许爬
func (p *BaseParser) sizeFilter(url string) bool {
	return true // 暂时取消size过滤
	res, err := httputil.DoGetWithTimeout(url, time.Second*1)
	if err != nil { // 有可能地址错误或网络不通
		return false
	}
	defer res.Body.Close()
	if res.ContentLength <= 0 {
		return true // 获取不到内容大小，假定为符合条件
	}

	if res.ContentLength > 1500000 { // 文件大于1.5MB（约）
		return false
	}
	return true
}

// true表示满足爬虫过滤条件，允许爬
func (p *BaseParser) Filter(url string) bool {
	if strutil.Len(url) < 15 { // 异常url
		return false
	}

	url = strings.ToLower(url)
	url = strings.TrimSpace(url)

	// 排除css,js等
	filters := []string{".css", ".js",
		".ico", ".jpg", ".jpeg", ".png", ".bmp", ".tif", ".gif", // 图片
		".mp3", ".asf", ".wma", ".wav", ".rm", ".real", ".ape", ".midi", ".flac", ".vqf", ".cd", ".ogg", // 音频
		".mp4", ".rm", ".rmvb", ".wmv", ".avi", ".3gp", ".mkv", ".flv", ".mpeg", ".mov", ".dat", ".f4v", ".mpg", // 视频
		".zip", ".7z", ".gz", ".rar", ".bz2", ".tar", ".iso", ".cab", ".xz", ".parcel", ".cbr", ".cbz", // 压缩文件
		".exe", ".pkg", ".rpm", ".deb", ".apk", ".ipa", ".dll", ".dmg", ".msi", ".cbr", ".ts", // 安装文件
		".txt", ".pdf", ".doc", ".docx", ".ppt", ".pptx", ".xls", ".xlsx", ".wps", ".log", ".epub", ".json", ".epub", ".chm", // 文档文件
		".bin", ".bak", ".cue", ".pet", ".mdf", ".mds", ".mdx", ".cbz", ".520", ".ace", ".csf",
		".daa", ".divx", ".dts", ".gho", ".img", ".ipa", ".isz", ".ivc", ".jar", ".ts"} // 其它文件
	for _, v := range filters {
		if strings.HasSuffix(url, v) {
			return false
		}
	}

	// 包含如下字段表示不是合法的url
	strs := []string{"javascript:"}
	for _, v := range strs {
		if strings.Contains(url, v) {
			return false
		}
	}

	// 非http和https开头的非法
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return false
	}

	// 页面大小限制
	if !p.sizeFilter(url) {
		return false
	}

	return true
}

func (p *BaseParser) Parse(doc *goquery.Document) []string {
	var urls = make([]string, 0, 1000)
	if doc == nil {
		return urls
	}

	// 爬所有链接
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		v, _ := s.Attr("href")
		urls = append(urls, v)
	})
	return urls
}

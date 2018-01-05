package parser

import (
	"gocrawler/db"
	"strings"
	"gocrawler/util/strutil"
	"net/http"
)

type BaseParser struct {
	startUrl string
	host     string
	id       string
	//urlQueue  *data.UrlQueue
	myDB *db.AppDB
}

func (p *BaseParser) GetStartUrl() string {
	return p.startUrl
}

func (p *BaseParser) GetHost() string {
	return p.host
}

func (p *BaseParser) GetId() string {
	return p.id
}

// true表示满足爬虫过滤条件，允许爬
func (p *BaseParser) sizeFilter(url string) bool {
	return true
	res, err := http.Get(url)
	if err != nil { // 有可能地址错误或网络不通
		return false
	}
	if res.ContentLength <= 0 {
		return true // 获取不到内容大小，假定为符合条件
	}

	if res.ContentLength > 2000000 { // 文件大于3MB（约）
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
		".ico", ".jpg", ".jpeg", ".png", ".bmp", ".tif", ".gif",                                         // 图片
		".mp3", ".asf", ".wma", ".wav", ".rm", ".real", ".ape", ".midi", ".flac", ".vqf", ".cd", ".ogg", // 音频
		".mp4", ".rm", ".rmvb", ".wmv", ".avi", ".3gp", ".mkv", ".flv", ".mpeg",                         // 视频
		".zip", ".7z", ".gz", ".rar", ".bz2", ".tar", ".iso", ".cab", ".xz", ".parcel",                  // 压缩文件
		".exe", ".pkg", ".rpm", ".deb", ".apk", ".ipa", ".dll", ".dmg", ".msi",                          // 安装文件
		".txt", ".pdf", ".doc", ".docx", ".ppt", ".xls", ".xlsx", ".wps", ".log", ".epub", ".json",      // 文档文件
		".bin", ".bak", "javascript:;"} // 其它文件
	for _, v := range filters {
		if strings.HasSuffix(url, v) {
			return false
		}
	}

	if !p.sizeFilter(url) {
		return false
	}

	return true
}

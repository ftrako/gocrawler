package bean

type AppBean struct {
	AppId        string // 格式一般为com.xxx.xxx
	StoreId      string // 商店ID，比如豌豆荚id为wandoujia，ios为apple
	IosAppId     string // ios专用的appid，一般为数字
	Name         string
	Category     string // 格式 ;分类1;分类2;
	Version      string
	MinVersion   string
	Os           string // android or ios
	Vender       string // 开发商
	Size         string
	UpdateTime   string // 更新时间
	InstallCount string // 安装次数
}

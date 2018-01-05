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

// CREATE TABLE `app` (
// 	`id` VARCHAR(50) NOT NULL COMMENT '格式：md5(appid+os)',
// 	`appid` MEDIUMTEXT NULL COMMENT '一般格式为com.xx.xxx',
// 	`storeid` MEDIUMTEXT NULL COMMENT '应用商店id',
// 	`ios_appid` MEDIUMTEXT NULL COMMENT 'ios中自动分配的appid，一般为数字',
// 	`name` MEDIUMTEXT NULL,
// 	`category` MEDIUMTEXT NULL COMMENT '分类，格式 ;分类1;分类2;',
// 	`version` MEDIUMTEXT NULL,
// 	`minversion` MEDIUMTEXT NULL,
// 	`os` MEDIUMTEXT NULL COMMENT 'android or ios',
// 	`vender` MEDIUMTEXT NULL COMMENT '开发商',
// 	`size` MEDIUMTEXT NULL,
// 	`update_time` MEDIUMTEXT NULL COMMENT '更新时间',
// 	`install_count` MEDIUMTEXT NULL COMMENT '安装次数',
// 	PRIMARY KEY (`id`)
// )
// COLLATE='utf8_general_ci'
// ENGINE=InnoDB
// ;

package bean

type SongBean struct {
	Name     string
	Singer   string
	Album    string // 专辑
	Size     string // 大小
	Date     string // 日期
	Language string // 语言类别
	Type     string // 类型，比如mp3，ape，flac等
	Url      string // 链接地址
	Download string // 下载地址
	Code     string // 提取码
}

//CREATE TABLE `song` (
//`id` CHAR(50) NOT NULL COMMENT 'md5(name+singer+album+type)',
//`name` CHAR(100) NULL DEFAULT NULL COMMENT '歌名',
//`singer` CHAR(30) NULL DEFAULT NULL COMMENT '演唱者',
//`album` CHAR(100) NULL DEFAULT NULL COMMENT '专辑',
//`size` CHAR(20) NULL DEFAULT NULL COMMENT '大小',
//`date` CHAR(20) NULL DEFAULT NULL COMMENT '日期',
//`language` CHAR(20) NULL DEFAULT NULL COMMENT '语言类别：国语，英语等',
//`type` CHAR(20) NULL DEFAULT NULL COMMENT '类型，比如mp3，ape，flac等',
//`url` TEXT NULL COMMENT '歌曲链接',
//`download` TEXT NULL COMMENT '下载地址',
//`code` CHAR(20) NULL DEFAULT NULL COMMENT '提取码',
//PRIMARY KEY (`id`)
//)
//COLLATE='utf8_general_ci'
//ENGINE=InnoDB
//;

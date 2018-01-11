package bean

type FileBean struct {
	Name       string // 文件名
	Suffix     string // 后缀名
	Url        string // 当前网站地址
	Download   string // 下载地址
	Pwd        string // 提取密码
	UnzipPwd   string // 解压密码
	Type       string // 文件类型
	Size       string // 文件大小
	UpdateDate string // 更新日期
	Author     string // 作者
}

//CREATE TABLE `file` (
//`id` CHAR(50) NOT NULL COMMENT 'md5(url+download)',
//`name` VARCHAR(200) NULL DEFAULT NULL COMMENT '文件名',
//`suffix` VARCHAR(10) NULL DEFAULT NULL COMMENT '后缀名',
//`url` VARCHAR(300) NULL DEFAULT NULL COMMENT '当前网站地址',
//`download` VARCHAR(500) NULL DEFAULT NULL COMMENT '下载地址',
//`pwd` VARCHAR(50) NULL DEFAULT NULL COMMENT '提取密码',
//`unzip_pwd` VARCHAR(50) NULL DEFAULT NULL COMMENT '解压密码',
//`type` VARCHAR(50) NULL DEFAULT NULL COMMENT '类型：book，music，video，exe，dmg等',
//`size` VARCHAR(50) NULL DEFAULT NULL COMMENT '大小',
//`update_date` VARCHAR(50) NULL DEFAULT NULL COMMENT '更新日期',
//`author` VARCHAR(100) NULL DEFAULT NULL COMMENT '作者',
//PRIMARY KEY (`id`)
//)
//COMMENT='存储各类文件地址'
//COLLATE='utf8_general_ci'
//ENGINE=InnoDB
//;

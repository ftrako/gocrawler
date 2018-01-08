package bean

type FileBean struct {
	Name     string // 文件名
	Suffix   string // 后缀名
	Url      string // 当前网站地址
	Download string // 下载地址
	Pwd      string // 提取密码
	ZipPwd   string // 解压密码
	Type     string // 文件类型
}

//CREATE TABLE `file` (
//`id` CHAR(50) NOT NULL COMMENT 'md5(url+download)',
//`name` VARCHAR(50) NULL DEFAULT NULL COMMENT '文件名',
//`suffix` VARCHAR(10) NULL DEFAULT NULL COMMENT '后缀名',
//`url` VARCHAR(200) NULL DEFAULT NULL COMMENT '当前网站地址',
//`download` VARCHAR(200) NULL DEFAULT NULL COMMENT '下载地址',
//`pwd` VARCHAR(20) NULL DEFAULT NULL COMMENT '提取密码',
//`zip_pwd` VARCHAR(20) NULL DEFAULT NULL COMMENT '压缩密码',
//`type` CHAR(10) NULL DEFAULT NULL COMMENT '类型：book，music，video，exe，dmg等',
//PRIMARY KEY (`id`)
//)
//COMMENT='存储各类文件地址'
//COLLATE='utf8_general_ci'
//ENGINE=InnoDB
//;

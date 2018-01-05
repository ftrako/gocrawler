package bean

type CategoryBean struct {
	Name      string
	SuperName string
	StoreId   string // 应用商店id
	StoreName string // 应用商店名称
}

// CREATE TABLE `category` (
// 	`id` VARCHAR(50) NOT NULL COMMENT '格式：md5(name+supername+stroreid)',
// 	`name` MEDIUMTEXT NULL,
// 	`super_name` MEDIUMTEXT NULL,
// 	`storeid` MEDIUMTEXT NULL COMMENT '应用商店id',
// 	`store_name` MEDIUMTEXT NULL COMMENT '应用商店名称',
// 	PRIMARY KEY (`id`)
// )
// COLLATE='utf8_general_ci'
// ENGINE=InnoDB
// ;
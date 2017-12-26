package db

import (
	"database/sql"
	"fmt"
	"gocrawler/bean"
	"gocrawler/util/cryptutil"

	_ "github.com/go-sql-driver/mysql"
)

var mysqlDB *sql.DB

// 表创建语句 app表
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

// category表
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

func Open() {
	db, err := sql.Open("mysql", "root:@tcp(10.0.2.206:3306)/app?charset=utf8")
	checkError(err)
	mysqlDB = db
}

func Close() {
	if mysqlDB == nil {
		return
	}

	mysqlDB.Close()
}

func ReplaceApp(bean *bean.AppBean) {
	if bean == nil || bean.AppId == "" {
		return
	}

	stmt, err := mysqlDB.Prepare("replace into app values(?,?,?,?,?,?,?,?,?,?,?,?,?);")
	checkError(err)
	_, err2 := stmt.Exec(cryptutil.MD5(bean.AppId+bean.Os),
		bean.AppId,
		bean.StoreId,
		bean.IosAppId,
		bean.Name,
		bean.Category,
		bean.Version,
		bean.MinVersion,
		bean.Os,
		bean.Vender,
		bean.Size,
		bean.UpdateTime,
		bean.InstallCount)
	if stmt != nil {
		stmt.Close()
	}
	checkError(err2)

	if err2 != nil {
		fmt.Println(err2.Error())
	}
}

func ReplaceCategory(bean *bean.CategoryBean) {
	if bean == nil || bean.Name == "" {
		return
	}
	stmt, err := mysqlDB.Prepare("replace into category values(?,?,?,?,?);")
	checkError(err)
	_, err2 := stmt.Exec(cryptutil.MD5(bean.Name+bean.SuperName+bean.StoreId),
		bean.Name,
		bean.SuperName,
		bean.StoreId,
		bean.StoreName)
	if stmt != nil {
		stmt.Close()
	}
	checkError(err2)

	if err2 != nil {
		fmt.Println(err2.Error())
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

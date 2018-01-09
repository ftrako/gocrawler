package db

import (
	"fmt"
	"gocrawler/bean"
	"gocrawler/util/cryptutil"
)

type AppDB struct {
	DB
}

func NewAppDB() *AppDB {
	myDB := new(AppDB)
	myDB.Open("mysql", "root:@tcp(10.0.2.206:3306)/app?charset=utf8")
	return myDB
}

func (p *AppDB) ReplaceApp(bean *bean.AppBean) {
	if bean == nil || bean.AppId == "" {
		return
	}

	stmt, err := p.myDB.Prepare("replace into app values(?,?,?,?,?,?,?,?,?,?,?,?,?);")
	p.checkError(err)
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
	p.checkError(err2)

	if err2 != nil {
		fmt.Println(err2.Error())
	}
}

func (p *AppDB) ReplaceCategory(bean *bean.CategoryBean) {
	if bean == nil || bean.Name == "" {
		return
	}
	stmt, err := p.myDB.Prepare("replace into category values(?,?,?,?,?,?);")
	p.checkError(err)
	_, err2 := stmt.Exec(cryptutil.MD5(bean.Name+bean.SuperName+bean.StoreId),
		bean.Cid,
		bean.Name,
		bean.SuperName,
		bean.StoreId,
		bean.StoreName)
	if stmt != nil {
		stmt.Close()
	}
	p.checkError(err2)

	if err2 != nil {
		fmt.Println(err2.Error())
	}
}

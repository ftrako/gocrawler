package db

import (
	"gocrawler/bean"
	"gocrawler/util/cryptutil"
)

type FileDB struct {
	DB
}

func NewFileDB() *FileDB {
	myDB := new(FileDB)
	myDB.initialize()
	return myDB
}

func (p *FileDB) initialize() {
	p.Open("mysql", "root:@tcp(10.0.2.206:3306)/file?charset=utf8")
}

func (p *FileDB) ReplaceFile(bean *bean.FileBean) {
	defer p.insertLocker.Unlock()
	p.insertLocker.Lock()
	if bean == nil || bean.Name == "" || bean.Type == "" {
		return
	}

	stmt, err := p.myDB.Prepare("replace into file values(?,?,?,?,?,?,?,?,?,?,?);")
	p.checkError(err)
	_, err2 := stmt.Exec(cryptutil.MD5(bean.Url+bean.Download),
		bean.Name,
		bean.Suffix,
		bean.Url,
		bean.Download,
		bean.Pwd,
		bean.UnzipPwd,
		bean.Type,
		bean.Size,
		bean.UpdateDate,
		bean.Author)
	if stmt != nil {
		stmt.Close()
	}
	p.checkError(err2)
}

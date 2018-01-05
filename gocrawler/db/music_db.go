package db

import (
	"gocrawler/util/cryptutil"
	"gocrawler/bean"
)

type MusicDB struct {
	DB
}

func NewMusicDB() *MusicDB {
	myDB := new(MusicDB)
	myDB.Open("mysql", "root:@tcp(10.0.2.206:3306)/music?charset=utf8")
	return myDB
}

func (p *MusicDB) ReplaceSong(bean *bean.SongBean) {
	defer p.insertLocker.Unlock()
	p.insertLocker.Lock()
	if bean == nil || bean.Name == "" {
		return
	}

	stmt, err := p.myDB.Prepare("replace into song values(?,?,?,?,?,?,?,?,?,?,?);")
	p.checkError(err)
	_, err2 := stmt.Exec(cryptutil.MD5(bean.Name+bean.Singer+bean.Album+bean.Type),
		bean.Name,
		bean.Singer,
		bean.Album,
		bean.Size,
		bean.Date,
		bean.Language,
		bean.Type,
		bean.Url,
		bean.Download,
		bean.Code)
	if stmt != nil {
		stmt.Close()
	}
	p.checkError(err2)
}

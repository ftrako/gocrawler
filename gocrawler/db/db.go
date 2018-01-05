package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

type DB struct {
	myDB *sql.DB
	insertLocker sync.Mutex
}

func (p *DB) Open(driver string, dataSource string) error {
	var err error
	p.myDB, err = sql.Open(driver, dataSource)
	return err
}

func (p *DB) Close() {
	if p.myDB != nil {
		p.myDB.Close()
	}
}

func (p *DB) checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

package backup

import (
	"os"
	"encoding/gob"
	"sync"
	"gocrawler/util/fileutil"
)

type Backup struct {
	path   string
	locker sync.Mutex
}

func NewBackup(path string) *Backup {
	backup := new(Backup)
	backup.setupData(path)
	return backup
}

func (p *Backup) setupData(path string) {
	p.path = path

	// 自动递归创建目录
	fileutil.MakeAllDirs(path)
}

func (p *Backup) LoadData() map[string]string {
	p.locker.Lock()
	defer p.locker.Unlock()
	var data map[string]string
	f, err := os.Open(p.path)
	if err != nil {
		return data
	}
	defer f.Close()
	dec := gob.NewDecoder(f)
	dec.Decode(&data)
	return data
}

func (p *Backup) Backup(obj interface{}) {
	p.locker.Lock()
	defer p.locker.Unlock()

	//// 空数据返回
	//if v, ok := obj.(map[string]string); ok {
	//	if len(v) <= 0 {
	//		return
	//	}
	//}

	f, err := os.Create(p.path)
	if err != nil {
		return
	}
	defer f.Close()
	enc := gob.NewEncoder(f)
	enc.Encode(obj)
}

func (p *Backup) Close() {
}

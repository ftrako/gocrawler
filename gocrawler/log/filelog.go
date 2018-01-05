package log

import (
	"os"
	"gocrawler/util/fileutil"
)

type FileLog struct {
	file *os.File
}

func NewFileLog(path string) *FileLog {
	fl := new(FileLog)
	fl.SetupData(path)
	return fl
}

func (p *FileLog) SetupData(path string) {
	fileutil.MakeAllDirs(path)
	p.file, _ = os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
}

func (p *FileLog) Println(text string) {
	p.file.WriteString(text + "\n")
}

func (p *FileLog) Release() {
	if p.file != nil {
		p.file.Close()
	}
}

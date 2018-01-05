package fileutil

import (
	"os"
	"gocrawler/util/strutil"
)

func Exist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func MakeAllDirs(filePath string) {
	parentPath := GetParentDirPath(filePath)
	os.MkdirAll(parentPath, os.ModePerm)
}

func GetParentDirPath(filePath string) string {
	index := strutil.LastIndex(filePath, "/")
	if index < 0 {
		return filePath
	}
	return strutil.SubString(filePath, 0, index)
}

package cryptutil

import (
	"crypto/md5"
	"fmt"
)

// MD5 md5加密
func MD5(text string) string {
	data := []byte(text)
	has := md5.Sum(data)
	md5 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5
}

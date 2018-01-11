package strutil

import (
	"github.com/axgle/mahonia"
	"strings"
)

func SubString(str string, start int, end int) string {
	if str == "" {
		return ""
	}
	innerText := []rune(str)
	count := len(innerText)
	if start >= end || end > count || start < 0 {
		return ""
	}
	return string(innerText[start:end])
}

func Suffix(str string) string {
	index := LastIndex(str, ".")
	if index < 0 {
		return ""
	}
	return SubString(str, index+1, Len(str))
}

func Len(str string) int {
	return len([]rune(str))
}

func Index(str string, substr string) int {
	index := strings.Index(str, substr)
	if index < 0 {
		return -1
	}
	prefix := []byte(str)[0:index]
	rs := []rune(string(prefix))
	return len(rs)
}

func LastIndex(str string, substr string) int {
	index := strings.LastIndex(str, substr)
	if index < 0 {
		return -1
	}
	prefix := []byte(str)[0:index]
	rs := []rune(string(prefix))
	return len(rs)
}

func Gbk2Utf8(str string) string {
	srcCode := "gbk"
	tagCode := "utf-8"
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(str)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func Utf82Gbk(str string) string {
	srcCode := "utf-8"
	tagCode := "gbk"
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(str)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

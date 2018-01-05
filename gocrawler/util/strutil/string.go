package strutil

import "strings"

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

func Len(str string) int {
	return len([]rune(str))
}

func Index(str string, substr string) int {
	index:=strings.Index(str, substr)
	if index < 0 {
		return -1
	}
	prefix := []byte(str)[0:index]
	rs:=[]rune(string(prefix))
	return len(rs)
}

func LastIndex(str string, substr string) int {
	index:=strings.LastIndex(str, substr)
	if index < 0 {
		return -1
	}
	prefix := []byte(str)[0:index]
	rs:=[]rune(string(prefix))
	return len(rs)
}
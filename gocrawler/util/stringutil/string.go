package stringutil

func SubString(text string, start int, end int) string {
	if text == "" {
		return ""
	}
	innerText := []rune(text)
	count := len(innerText)
	if start >= end || end >= count || start < 0 {
		return ""
	}
	return string(innerText[start:end])
}

package iteration

import "strings"

func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

func IfContainsReturnString(s string, substr string) string {
	if strings.Contains(s, substr) {
		return s
	}
	return ""
}

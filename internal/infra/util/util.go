package util

import (
	"strings"
)

func PadText(str string, length int) string {
	if len(str) >= length {
		return str
	}
	return str + strings.Repeat(" ", length-len(str))
}

func Contains(str string, l []string) bool {
	for _, v := range l {
		if str == v {
			return true
		}
	}
	return false
}

func KeySlice[T comparable, U any](m map[T]U) []T {
	keys := make([]T, 0)
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func HighestLength(items []string) int {
	leng := 0
	for _, v := range items {
		if len(v) > leng {
			leng = len(v)
		}
	}
	return leng
}

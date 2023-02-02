package util

import "strings"

func InArray[T comparable](needle T, array []T) bool {
	for _, item := range array {
		if needle == item {
			return true
		}
	}
	return false
}

func ContainsArray(str string, array []string) bool {
	for _, item := range array {
		if strings.Contains(str, item) {
			return true
		}
	}
	return false
}

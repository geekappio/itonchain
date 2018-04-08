package util

import "strings"

/**
 *	是否包含substr字符串
 */
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// 判断字符串是否相同，忽略大小写
func EqualsIgnoreCase(s1, s2 string) bool {
	return strings.ToLower(s1) == strings.ToLower(s2)
}

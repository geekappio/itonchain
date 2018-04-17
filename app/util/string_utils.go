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

// Check if specified string is nil.
func StringIsNil(str string) bool {
	return str == ""
}

// Check if specified string is nil or is space.
func StringIsBlack(str string) bool {
	return StringIsNil(str) || strings.Trim(str, " ") == ""
}
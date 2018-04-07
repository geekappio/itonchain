package logging

import "strings"

/**
 *	是否包含substr字符串
 */
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}


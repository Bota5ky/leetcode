package repeated_substring_pattern

import "strings"

// https://leetcode.cn/problems/repeated-substring-pattern/
func repeatedSubstringPattern(s string) bool {
	return strings.Contains((s + s)[1:len(s+s)-1], s)
}

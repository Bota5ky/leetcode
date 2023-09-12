package repeated_substring_pattern

import "strings"

// 459. 重复的子字符串 https://leetcode.cn/problems/repeated-substring-pattern/
func repeatedSubstringPattern(s string) bool {
	return strings.Contains((s + s)[1:len(s+s)-1], s)
}

package longest_happy_prefix

import "strings"

// 1392. 最长快乐前缀 https://leetcode.cn/problems/longest-happy-prefix/
func longestPrefix(s string) string {
	prefix := s[1:]
	for strings.HasPrefix(s, prefix) == false {
		prefix = prefix[1:]
	}
	return prefix
}

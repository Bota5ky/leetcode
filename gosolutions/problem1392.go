package gosolutions

import "strings"

// https://leetcode-cn.com/problems/longest-happy-prefix/
func longestPrefix(s string) string {
	prefix := s[1:]
	for strings.HasPrefix(s, prefix) == false {
		prefix = prefix[1:]
	}
	return prefix
}

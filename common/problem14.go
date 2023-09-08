package common

import "strings"

// https://leetcode-cn.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs) && len(prefix) > 0; i++ {
		for len(prefix) > 0 && strings.HasPrefix(strs[i], prefix) == false {
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

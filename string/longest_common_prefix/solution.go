package longest_common_prefix

import "strings"

// 14. 最长公共前缀 https://leetcode.cn/problems/longest-common-prefix/
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

package string_to_url_lcci

import "strings"

// 面试题 01.03. URL化 https://leetcode.cn/problems/string-to-url-lcci/
func replaceSpaces(S string, length int) string {
	return strings.Replace(S[:length], " ", "%20", -1)
}

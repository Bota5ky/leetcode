package offer

import "strings"

// https://leetcode.cn/problems/ti-huan-kong-ge-lcof/
func replaceSpace(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}

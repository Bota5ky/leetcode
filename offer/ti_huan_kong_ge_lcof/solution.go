package offer

import "strings"

// 剑指 Offer 05. 替换空格 https://leetcode.cn/problems/ti-huan-kong-ge-lcof/
func replaceSpace(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}

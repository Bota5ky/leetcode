package gosolutions

import "strings"

// https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
func replaceSpace(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}

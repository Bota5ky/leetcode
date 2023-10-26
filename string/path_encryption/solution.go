package ti_huan_kong_ge_lcof

import "strings"

// LCR 122. 路径加密 https://leetcode.cn/problems/ti-huan-kong-ge-lcof/
func pathEncryption(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}

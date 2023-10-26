package zuo_xuan_zhuan_zi_fu_chuan_lcof

// LCR 182. 动态口令 https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/
func dynamicPassword(s string, n int) string {
	return s[n:] + s[:n]
}

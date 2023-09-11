package zuo_xuan_zhuan_zi_fu_chuan_lcof

// 剑指 Offer 58 - II. 左旋转字符串 https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/
func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}

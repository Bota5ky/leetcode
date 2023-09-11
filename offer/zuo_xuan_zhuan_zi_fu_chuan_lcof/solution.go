package offer

// https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/
func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}

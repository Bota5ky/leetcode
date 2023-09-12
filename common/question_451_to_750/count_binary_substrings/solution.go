package count_binary_substrings

// 696. 计数二进制子串 https://leetcode.cn/problems/count-binary-substrings/
func countBinarySubstrings(s string) int {
	res := 0
	pre := s[0]
	preCnt, curCnt := 0, 1
	for i := 1; i < len(s); i++ {
		if s[i] != pre {
			preCnt = curCnt
			curCnt = 1
		} else {
			curCnt++
		}
		pre = s[i]
		if preCnt >= curCnt {
			res++
		}
	}
	return res
}

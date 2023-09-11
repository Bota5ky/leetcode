package question_451_to_750

// https://leetcode.cn/problems/count-binary-substrings/solution/
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

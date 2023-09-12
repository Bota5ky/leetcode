package repeated_string_match

import "strings"

// 686. 重复叠加字符串匹配 https://leetcode.cn/problems/repeated-string-match/
func repeatedStringMatch(A string, B string) int {
	maxLen := len(A)*2 + len(B)
	cnt := 1
	C := A
	for len(C) <= maxLen {
		if strings.Contains(C, B) == false {
			C += A
			cnt++
		} else {
			return cnt
		}
	}
	return -1
}

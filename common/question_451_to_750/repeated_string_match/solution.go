package repeated_string_match

import "strings"

// https://leetcode.cn/problems/repeated-string-match/
func repeatedStringMatch(A string, B string) int {
	maxlen := len(A)*2 + len(B)
	cnt := 1
	C := A
	for len(C) <= maxlen {
		if strings.Contains(C, B) == false {
			C += A
			cnt++
		} else {
			return cnt
		}
	}
	return -1
}

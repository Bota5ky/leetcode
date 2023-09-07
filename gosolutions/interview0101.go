package gosolutions

import "strings"

// 面试题 01.01. 判定字符是否唯一 https://leetcode.cn/problems/is-unique-lcci/
func isUnique(astr string) bool {
	for i := 0; i < len(astr); i++ {
		if strings.Contains(astr[i+1:], string(astr[i])) {
			return false
		}
	}
	return true
}

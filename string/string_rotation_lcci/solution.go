package string_rotation_lcci

import "strings"

// 面试题 01.09. 字符串轮转 https://leetcode.cn/problems/string-rotation-lcci/
func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	temp := s1 + s1
	return strings.Contains(temp, s2)
}

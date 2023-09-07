package gosolutions

import "strings"

// https://leetcode-cn.com/problems/string-rotation-lcci/
func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	temp := s1 + s1
	return strings.Contains(temp, s2)
}

package leetcode

import "strings"

//https://leetcode-cn.com/problems/is-unique-lcci/
func isUnique(astr string) bool {
	for i := 0; i < len(astr); i++ {
		if strings.Contains(astr[i+1:], string(astr[i])) {
			return false
		}
	}
	return true
}

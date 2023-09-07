package leetcode

import "strings"

//https://leetcode-cn.com/problems/string-to-url-lcci/
func replaceSpaces(S string, length int) string {
	return strings.Replace(S[:length], " ", "%20", -1)
}

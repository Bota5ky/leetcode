package gosolutions

import "strings"

// https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/
func firstUniqChar(s string) byte {
	for i := 0; i < len(s); i++ {
		if strings.Count(s, string(s[i])) == 1 {
			return s[i]
		}
	}
	return ' '
}

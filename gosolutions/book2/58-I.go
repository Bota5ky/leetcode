package leetcode

import "strings"

//https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof/
func reverseWords(s string) string {
	var ret string
	words := strings.Split(s, " ")
	for i := 0; i < len(words); i++ {
		if words[i] != "" {
			ret = " " + words[i] + ret
		}
	}
	if len(ret) > 0 {
		ret = ret[1:]
	}
	return ret
}

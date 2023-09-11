package fan_zhuan_dan_ci_shun_xu_lcof

import "strings"

// 剑指 Offer 58 - I. 翻转单词顺序 https://leetcode.cn/problems/fan-zhuan-dan-ci-shun-xu-lcof/
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

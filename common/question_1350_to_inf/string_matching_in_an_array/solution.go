package string_matching_in_an_array

import "strings"

// https://leetcode.cn/problems/string-matching-in-an-array/
func stringMatching(words []string) []string {
	var res []string
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i == j || len(words[j]) <= len(words[i]) {
				continue
			}
			if strings.Contains(words[j], words[i]) {
				res = append(res, words[i])
				break
			}
		}
	}
	return res
}

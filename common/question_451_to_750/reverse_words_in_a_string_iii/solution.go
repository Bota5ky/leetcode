package reverse_words_in_a_string_iii

import "strings"

// https://leetcode.cn/problems/reverse-words-in-a-string-iii/
func reverseWords(s string) string {
	ts := []byte(s)
	for i, j := 0, len(ts)-1; i < j; {
		ts[i], ts[j] = ts[j], ts[i]
		i++
		j--
	}
	words := strings.Split(string(ts), " ")
	var ret string
	for i := len(words) - 1; i >= 0; i-- {
		ret += " " + words[i]
	}
	return ret[1:]
}

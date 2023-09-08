package common

import "strings"

// https://leetcode-cn.com/problems/valid-palindrome/
func isPalindrome2(s string) bool {
	str := strings.ToLower(s)
	i, j := 0, len(str)-1
	for i < j {
		for !(str[i] >= 'a' && str[i] <= 'z' || str[i] >= '0' && str[i] <= '9') && i < j {
			i++
		}
		for !(str[j] >= 'a' && str[j] <= 'z' || str[j] >= '0' && str[j] <= '9') && i < j {
			j--
		}
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}

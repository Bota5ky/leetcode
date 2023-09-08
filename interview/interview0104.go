package interview

// 面试题 01.04. 回文排列 https://leetcode.cn/problems/palindrome-permutation-lcci/
func canPermutePalindrome(s string) bool {
	numOfOdd := 0
	m := make(map[rune]bool)
	for _, val := range s {
		if m[val] == false {
			m[val] = true
			numOfOdd++
		} else {
			m[val] = false
			numOfOdd--
		}
	}
	return numOfOdd <= 1
}

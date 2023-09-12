package valid_palindrome_ii

// https://leetcode.cn/problems/valid-palindrome-ii/
func validPalindrome(s string) bool {
	j := len(s) - 1
	for i := 0; i < j; i++ {
		if s[i] != s[j] {
			if isR(s[i+1:j+1]) || isR(s[i:j]) {
				return true
			}
			if isR(s[i+1:j+1]) == false && isR(s[i:j]) == false {
				return false
			}
		}
		j--
	}
	return true
}
func isR(s string) bool {
	j := len(s) - 1
	for i := 0; i < j; i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}
	return true
}

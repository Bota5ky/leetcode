package question_1001_to_1350

// https://leetcode.cn/problems/remove-palindromic-subsequences/
func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	}
	j := len(s) - 1
	for i := 0; i < j; i++ {
		if s[i] != s[j] {
			return 2
		}
		j--
	}
	return 1
}

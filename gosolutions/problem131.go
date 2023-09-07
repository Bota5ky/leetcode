package gosolutions

// https://leetcode-cn.com/problems/palindrome-partitioning/
func partition(s string) [][]string {
	var res [][]string
	for i := len(s) - 1; i >= 0; i-- {
		var temp [][]string
		if isPalindrome3(s[i:]) {
			temp = partition(s[:i]) //[][]string
			if i == 0 {
				res = append(res, []string{s[i:]})
			} else {
				for j := range temp {
					res = append(res, append(temp[j], s[i:]))
				}
			}
		}
	}
	return res
}
func isPalindrome3(s string) bool {
	for i, j := 0, len(s)-1; i < j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

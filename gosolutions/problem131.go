package gosolutions

// https://leetcode.cn/problems/palindrome-partitioning/
func partitionString(s string) [][]string {
	var res [][]string
	for i := len(s) - 1; i >= 0; i-- {
		var temp [][]string
		if isPalindrome3(s[i:]) {
			temp = partitionString(s[:i]) //[][]string
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

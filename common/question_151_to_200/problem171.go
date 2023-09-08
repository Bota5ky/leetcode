package question_151_to_200

// https://leetcode-cn.com/problems/excel-sheet-column-number/submissions/
func titleToNumber(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res = res*26 + int(s[i]-'A'+1)
	}
	return res
}

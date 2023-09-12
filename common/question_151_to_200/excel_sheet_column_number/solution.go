package excel_sheet_column_number

// https://leetcode.cn/problems/excel-sheet-column-number/submissions/
func titleToNumber(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res = res*26 + int(s[i]-'A'+1)
	}
	return res
}

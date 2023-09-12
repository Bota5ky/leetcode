package excel_sheet_column_number

// 171. Excel 表列序号 https://leetcode.cn/problems/excel-sheet-column-number/
func titleToNumber(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res = res*26 + int(s[i]-'A'+1)
	}
	return res
}

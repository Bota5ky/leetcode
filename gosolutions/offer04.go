package gosolutions

// https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m := 0
	n := len(matrix[0]) - 1
	for n >= 0 && m < len(matrix) {
		if matrix[m][n] == target {
			return true
		}
		for n >= 0 && matrix[m][n] > target {
			n--
		}
		for n >= 0 && m < len(matrix) && matrix[m][n] < target {
			m++
		}
	}
	return false
}

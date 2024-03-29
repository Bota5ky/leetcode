package lucky_numbers_in_a_matrix

// 1380. 矩阵中的幸运数 https://leetcode.cn/problems/lucky-numbers-in-a-matrix/
func luckyNumbers(matrix [][]int) []int {
	m := make(map[int]bool)
	var ret []int
	for i := 0; i < len(matrix); i++ {
		small := matrix[i][0]
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] < small {
				small = matrix[i][j]
			}
		}
		m[small] = true
	}
	for j := 0; j < len(matrix[0]); j++ {
		maximum := matrix[0][j]
		for i := 1; i < len(matrix); i++ {
			if matrix[i][j] > maximum {
				maximum = matrix[i][j]
			}
		}
		if m[maximum] {
			ret = append(ret, maximum)
		}
	}
	return ret
}

package common

// https://leetcode-cn.com/problems/lucky-numbers-in-a-matrix/
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
		max := matrix[0][j]
		for i := 1; i < len(matrix); i++ {
			if matrix[i][j] > max {
				max = matrix[i][j]
			}
		}
		if m[max] {
			ret = append(ret, max)
		}
	}
	return ret
}

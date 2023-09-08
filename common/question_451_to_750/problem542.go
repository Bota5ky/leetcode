package question_451_to_750

// https://leetcode-cn.com/problems/01-matrix/
func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]int{}
	}
	res := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		res[i] = make([]int, len(matrix[i]))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				res[i][j] = 0
			} else {
				res[i][j] = 1<<31 - 1
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i > 0 {
				res[i][j] = min(res[i][j], res[i-1][j]+1)
			}
			if j > 0 {
				res[i][j] = min(res[i][j], res[i][j-1]+1)
			}
		}
	}
	for i := len(matrix) - 1; i >= 0; i-- {
		for j := len(matrix[i]) - 1; j >= 0; j-- {
			if i < len(matrix)-1 {
				res[i][j] = min(res[i][j], res[i+1][j]+1)
			}
			if j < len(matrix[i])-1 {
				res[i][j] = min(res[i][j], res[i][j+1]+1)
			}
		}
	}
	return res
}

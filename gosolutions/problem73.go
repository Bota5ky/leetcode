package gosolutions

// https://leetcode-cn.com/problems/set-matrix-zeroes/
func setZeroes(matrix [][]int) {
	rec := 1 //代表第一列是否归零
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			rec = 0
			break
		}
	}
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			matrix[0][0] = 0 //代表第一行是否归零
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < len(matrix[i]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			for j := 1; j < len(matrix); j++ {
				matrix[j][i] = 0
			}
		}
	}
	if matrix[0][0] == 0 {
		for i := 1; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}
	if rec == 0 {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}

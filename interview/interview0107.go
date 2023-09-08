package interview

// 面试题 01.07. 旋转矩阵 https://leetcode.cn/problems/rotate-matrix-lcci/
func rotateMatrix(matrix [][]int) {
	//转置
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	//镜像
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i])/2; j++ {
			matrix[i][j], matrix[i][len(matrix[i])-1-j] = matrix[i][len(matrix[i])-1-j], matrix[i][j]
		}
	}
}

package spiral_matrix_ii

// 59. 螺旋矩阵 II https://leetcode.cn/problems/spiral-matrix-ii/
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	up, down, left, right := 0, n-1, 0, n-1
	for k := 1; k <= n*n; {
		//-->
		for i := left; i <= right && up <= down; i++ {
			res[up][i] = k
			k++
		}
		up++
		//↓
		for i := up; i <= down && left <= right; i++ {
			res[i][right] = k
			k++
		}
		right--
		//<--
		for i := right; i >= left && up <= down; i-- {
			res[down][i] = k
			k++
		}
		down--
		//↑
		for i := down; i >= up && left <= right; i-- {
			res[i][left] = k
			k++
		}
		left++
	}
	return res
}

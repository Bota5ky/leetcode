package spiral_matrix

// 和offer29相同
// https://leetcode.cn/problems/spiral-matrix/
func spiralOrder(matrix [][]int) []int {
	var res []int
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}
	up, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for up <= down && left <= right {
		//-->
		for i := left; i <= right && up <= down; i++ {
			res = append(res, matrix[up][i])
		}
		up++
		//↓
		for i := up; i <= down && left <= right; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		//<--
		for i := right; i >= left && up <= down; i-- {
			res = append(res, matrix[down][i])
		}
		down--
		//↑
		for i := down; i >= up && left <= right; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}

	return res
}

package question_301_to_450

// https://leetcode.cn/problems/longest-increasing-path-in-a-matrix/
func longestIncreasingPath(matrix [][]int) int {
	max := 0
	rec := make(map[[2]int]int)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			cnt := dfs2(matrix, i, j, -1<<31, &rec)
			if max < cnt {
				max = cnt
			}
		}
	}
	return max
}
func dfs2(matrix [][]int, i, j, preVal int, rec *map[[2]int]int) int {
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[i]) {
		return 0
	}
	if preVal >= matrix[i][j] {
		return 0
	}
	if (*rec)[[2]int{i - 1, j}] <= 0 {
		(*rec)[[2]int{i - 1, j}] = dfs2(matrix, i-1, j, matrix[i][j], rec)
	}
	if (*rec)[[2]int{i + 1, j}] <= 0 {
		(*rec)[[2]int{i + 1, j}] = dfs2(matrix, i+1, j, matrix[i][j], rec)
	}
	if (*rec)[[2]int{i, j - 1}] <= 0 {
		(*rec)[[2]int{i, j - 1}] = dfs2(matrix, i, j-1, matrix[i][j], rec)
	}
	if (*rec)[[2]int{i, j + 1}] <= 0 {
		(*rec)[[2]int{i, j + 1}] = dfs2(matrix, i, j+1, matrix[i][j], rec)
	}
	a := (*rec)[[2]int{i - 1, j}]
	if i > 0 && matrix[i-1][j] == matrix[i][j] {
		a = 0
	}
	b := (*rec)[[2]int{i + 1, j}]
	if i < len(matrix)-1 && matrix[i+1][j] == matrix[i][j] {
		b = 0
	}
	c := (*rec)[[2]int{i, j - 1}]
	if j > 0 && matrix[i][j-1] == matrix[i][j] {
		c = 0
	}
	d := (*rec)[[2]int{i, j + 1}]
	if j < len(matrix[i])-1 && matrix[i][j+1] == matrix[i][j] {
		d = 0
	}
	return 1 + m(a, b, c, d)
}
func m(a, b, c, d int) int {
	temp := a
	if temp < b {
		temp = b
	}
	if temp < c {
		temp = c
	}
	if temp < d {
		temp = d
	}
	return temp
}

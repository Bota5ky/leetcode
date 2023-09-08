package common

// https://leetcode-cn.com/problems/triangle/
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	var min int
	for i := 1; i < len(triangle); i++ {
		triangle[i][0] += triangle[i-1][0]
		l := len(triangle[i])
		triangle[i][l-1] += triangle[i-1][l-2]
		if i == len(triangle)-1 {
			if triangle[i][l-1] > triangle[i][0] {
				min = triangle[i][0]
			} else {
				min = triangle[i][l-1]
			}
		}
		for j := 1; j < l-1; j++ {
			if triangle[i-1][j-1] > triangle[i-1][j] {
				triangle[i][j] += triangle[i-1][j]
			} else {
				triangle[i][j] += triangle[i-1][j-1]
			}
			if i == len(triangle)-1 && triangle[i][j] < min {
				min = triangle[i][j]
			}
		}
	}
	return min
}

package question_751_to_1000

// https://leetcode.cn/problems/surface-area-of-3d-shapes/
func surfaceArea(grid [][]int) int {
	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 0 {
				sum += 2
			}
			temp := 0
			if i < len(grid)-1 {
				temp = grid[i+1][j]
			}
			sum += max(grid[i][j]-temp, 0)
			temp = 0
			if i > 0 {
				temp = grid[i-1][j]
			}
			sum += max(grid[i][j]-temp, 0)
			temp = 0
			if j < len(grid[i])-1 {
				temp = grid[i][j+1]
			}
			sum += max(grid[i][j]-temp, 0)
			temp = 0
			if j > 0 {
				temp = grid[i][j-1]
			}
			sum += max(grid[i][j]-temp, 0)
		}
	}
	return sum
}

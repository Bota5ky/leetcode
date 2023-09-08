package question_1001_to_1350

// https://leetcode-cn.com/problems/as-far-from-land-as-possible/
func maxDistance(grid [][]int) int {
	level := 1
	for {
		cnt := 0
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] == level {
					cnt += deal2(grid, i, j, level)
				}
			}
		}
		if cnt == 0 {
			break
		}
		level++
	}
	if level == 1 {
		return -1
	}
	return level - 1
}
func deal2(grid [][]int, i, j, level int) int {
	cnt := 0
	grid[i][j] = level + 1
	if i > 0 && grid[i-1][j] == 0 {
		cnt++
		grid[i-1][j] = level + 1
	}
	if i < len(grid)-1 && grid[i+1][j] == 0 {
		cnt++
		grid[i+1][j] = level + 1
	}
	if j > 0 && grid[i][j-1] == 0 {
		cnt++
		grid[i][j-1] = level + 1
	}
	if j < len(grid[i])-1 && grid[i][j+1] == 0 {
		cnt++
		grid[i][j+1] = level + 1
	}
	return cnt
}

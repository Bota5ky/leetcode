package offer

// https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/
func maxValue(grid [][]int) int {
	for i := 0; i < len(grid); i++ { //行数
		for j := 0; j < len(grid[i]); j++ { //列数
			if i == 0 {
				if j == 0 {
					continue
				}
				grid[i][j] += grid[i][j-1]
				continue
			}
			if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				if grid[i-1][j] > grid[i][j-1] {
					grid[i][j] += grid[i-1][j]
				} else {
					grid[i][j] += grid[i][j-1]
				}
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

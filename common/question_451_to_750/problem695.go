package question_451_to_750

// https://leetcode-cn.com/problems/max-area-of-island/
func maxAreaOfIsland(grid [][]int) int {
	//访问过的设置为0
	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			temp := area(grid, i, j)
			if max < temp {
				max = temp
			}
		}
	}
	return max
}

func area(grid [][]int, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	return 1 + area(grid, i-1, j) + area(grid, i+1, j) + area(grid, i, j-1) + area(grid, i, j+1)
}

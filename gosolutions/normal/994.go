package leetcode

//https://leetcode-cn.com/problems/rotting-oranges/
func orangesRotting(grid [][]int) int {
	time := 0
	for {
		infect := 0
		var temp [][2]int
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] == 2 {
					if i > 0 && grid[i-1][j] == 1 {
						temp = append(temp, [2]int{i - 1, j})
						infect++
					}
					if i < len(grid)-1 && grid[i+1][j] == 1 {
						temp = append(temp, [2]int{i + 1, j})
						infect++
					}
					if j > 0 && grid[i][j-1] == 1 {
						temp = append(temp, [2]int{i, j - 1})
						infect++
					}
					if j < len(grid[i])-1 && grid[i][j+1] == 1 {
						temp = append(temp, [2]int{i, j + 1})
						infect++
					}
				}
			}
		}
		if infect == 0 {
			break
		}
		time++
		for i := 0; i < len(temp); i++ {
			grid[temp[i][0]][temp[i][1]] = 2
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return time
}

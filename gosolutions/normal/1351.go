package leetcode

//https://leetcode-cn.com/problems/count-negative-numbers-in-a-sorted-matrix/
func countNegatives(grid [][]int) int {
	tag, cnt := len(grid[0]), 0
	for i := 0; i < len(grid); i++ {
		if tag == 0 {
			break
		}
		for j := 0; j < tag; j++ {
			if grid[i][j] >= 0 {
				cnt++
			} else {
				tag = j
				continue
			}
		}
	}
	return len(grid)*len(grid[0]) - cnt
}

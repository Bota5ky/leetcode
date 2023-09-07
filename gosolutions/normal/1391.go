package leetcode

//https://leetcode-cn.com/problems/check-if-there-is-a-valid-path-in-a-grid/
func hasValidPath(grid [][]int) bool {
	return dfsPath(grid, 0, 0, 1)
}
func dfsPath(grid [][]int, i, j, preDir int) bool {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
		return false
	}
	if grid[i][j] == 0 {
		return false
	} //是否遍历过
	temp := grid[i][j]
	grid[i][j] = 0
	//1.判断是否通行
	if i != 0 || j != 0 {
		switch preDir {
		case 1:
			if temp != 2 && temp != 3 && temp != 4 {
				return false
			} //上
		case 2:
			if temp != 2 && temp != 5 && temp != 6 {
				return false
			} //下
		case 3:
			if temp != 1 && temp != 4 && temp != 6 {
				return false
			} //左
		case 4:
			if temp != 1 && temp != 3 && temp != 5 {
				return false
			} //右
		}
	}
	//2.是否是终点
	if i == len(grid)-1 && j == len(grid[i])-1 {
		return true
	}
	//3.下一节点
	switch temp {
	case 1:
		return dfsPath(grid, i, j+1, 4) || dfsPath(grid, i, j-1, 3)
	case 2:
		return dfsPath(grid, i-1, j, 1) || dfsPath(grid, i+1, j, 2)
	case 3:
		return dfsPath(grid, i+1, j, 2) || dfsPath(grid, i, j-1, 3)
	case 4:
		return dfsPath(grid, i+1, j, 2) || dfsPath(grid, i, j+1, 4)
	case 5:
		return dfsPath(grid, i-1, j, 1) || dfsPath(grid, i, j-1, 3)
	case 6:
		return dfsPath(grid, i-1, j, 1) || dfsPath(grid, i, j+1, 4)
	}
	return true
}

package robot_in_a_grid_lcci

// https://leetcode-cn.com/problems/robot-in-a-grid-lcci/
func pathWithObstacles(obstacleGrid [][]int) [][]int {
	var res, routine [][]int
	visited := make([][]bool, len(obstacleGrid))
	for i := 0; i < len(obstacleGrid); i++ {
		visited[i] = make([]bool, len(obstacleGrid[i]))
	}
	dfs3(obstacleGrid, 0, 0, routine, &visited, &res)
	return res
}

// 保存之前的路径
// 终点是右下角 或者遇到障碍点 返回
func dfs3(obstacleGrid [][]int, r int, c int, routine [][]int, visited *([][]bool), res *([][]int)) {
	if obstacleGrid[r][c] == 1 || (*visited)[r][c] {
		return
	}
	routine = append(routine, []int{r, c})
	(*visited)[r][c] = true
	//判断是否是终点
	if r == len(obstacleGrid)-1 && c == len(obstacleGrid[r])-1 {
		for i := 0; i < len(routine); i++ {
			*res = append(*res, routine[i])
		}
		return
	}
	if r < len(obstacleGrid)-1 {
		dfs3(obstacleGrid, r+1, c, routine, visited, res)
	}
	if c < len(obstacleGrid[r])-1 && len(*res) == 0 {
		dfs3(obstacleGrid, r, c+1, routine, visited, res)
	}
}

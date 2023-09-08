package question_51_to_100

// https://leetcode-cn.com/problems/unique-paths-ii/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	cho := make([][]int, len(obstacleGrid))
	for i := 0; i < len(cho); i++ {
		cho[i] = make([]int, len(obstacleGrid[0]))
	}
	if obstacleGrid[0][0] == 0 {
		cho[0][0] = 1
	}
	for i := 0; i < len(cho); i++ {
		for j := 0; j < len(cho[i]); j++ {
			helper1(cho, obstacleGrid, i, j)
		}
	}
	return cho[len(cho)-1][len(cho[0])-1]
}
func helper1(cho, obstacleGrid [][]int, i, j int) {
	if obstacleGrid[i][j] == 1 {
		return
	}
	if i > 0 && obstacleGrid[i-1][j] != 1 {
		cho[i][j] += cho[i-1][j]
	}
	if j > 0 && obstacleGrid[i][j-1] != 1 {
		cho[i][j] += cho[i][j-1]
	}
}

package leetcode

//https://leetcode-cn.com/problems/magic-squares-in-grid/
func numMagicSquaresInside(grid [][]int) int {
	cnt := 0
	for i := 0; i < len(grid)-2; i++ { //行号
		for j := 0; j < len(grid[i])-2; j++ { //列号
			if isMagic(grid, i, j) {
				cnt++
			}
		}
	}
	return cnt
}
func isMagic(grid [][]int, a int, b int) bool {
	m:=make(map[int]int,9)
	for i := a; i <= a+2; i++ {
		for j := b; j <= b+2; j++ {
			if grid[i][j]<1 || grid[i][j]>9 {return false}
			if m[grid[i][j]]>0 {return false}
		}
	}
	if grid[a][b]+grid[a][b+1]+grid[a][b+2]!=15 {return false}
	if grid[a+1][b]+grid[a+1][b+1]+grid[a+1][b+2]!=15 {return false}
	if grid[a+2][b]+grid[a+2][b+1]+grid[a+2][b+2]!=15 {return false}
	if grid[a][b]+grid[a+1][b]+grid[a+2][b]!=15 {return false}
	if grid[a][b+1]+grid[a+1][b+1]+grid[a+2][b+1]!=15 {return false}
	if grid[a][b+2]+grid[a+1][b+2]+grid[a+2][b+2]!=15 {return false}
	if grid[a][b]+grid[a+1][b+1]+grid[a+2][b+2]!=15 {return false}
	if grid[a+2][b]+grid[a+1][b+1]+grid[a][b+2]!=15 {return false}
	return true
}

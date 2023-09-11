package question_751_to_1000

// https://leetcode.cn/problems/available-captures-for-rook/
func numRookCaptures(board [][]byte) int {
	dir := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	//find the ROOK
	x, y := -1, -1
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'R' {
				x, y = i, j
			}
		}
	}
	cnt := 0
	for i := 0; i < 4; i++ {
		m, n := x, y
		for {
			m += dir[i][0]
			n += dir[i][1]
			if m < 0 || m >= len(board) || n < 0 || n >= len(board[m]) {
				break
			}
			if board[m][n] == 'B' {
				break
			}
			if board[m][n] == 'p' {
				cnt++
				break
			}
		}
	}
	return cnt
}

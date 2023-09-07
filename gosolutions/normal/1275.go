package leetcode

//https://leetcode-cn.com/problems/find-winner-on-a-tic-tac-toe-game/
func tictactoe(moves [][]int) string {
	var square [3][3]byte
	for i := 0; i < 4 && i < len(moves); i++ {
		if i%2 == 1 {
			square[moves[i][0]][moves[i][1]] = 'O'
		} else {
			square[moves[i][0]][moves[i][1]] = 'X'
		}
	}
	for i := 4; i < len(moves); i++ {
		x := moves[i][0]
		y := moves[i][1]
		if i%2 == 1 {
			square[x][y] = 'O'
		} else {
			square[x][y] = 'X'
		}
		if isThreesome(square, x, y) {
			var ret string
			if i%2 == 0 {
				ret = "A"
			} else {
				ret = "B"
			}
			return ret
		}
		if i == 8 {
			return "Draw"
		}
	}
	return "Pending"
}
func isThreesome(square [3][3]byte, x int, y int) bool {
	if square[x][0] == square[x][1] && square[x][0] == square[x][2] {
		return true
	}
	if square[0][y] == square[1][y] && square[0][y] == square[2][y] {
		return true
	}
	if square[0][0] != 0 && square[0][0] == square[1][1] && square[0][0] == square[2][2] {
		return true
	}
	if square[1][1] != 0 && square[2][0] == square[1][1] && square[1][1] == square[0][2] {
		return true
	}
	return false
}

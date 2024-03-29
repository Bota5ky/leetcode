package game_of_life

// 289. 生命游戏 https://leetcode.cn/problems/game-of-life/
func gameOfLife(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			cnt := numOfLiveCell(board, i, j)
			if board[i][j] == 0 {
				if cnt == 3 {
					board[i][j] = -1
				}
			} else {
				if cnt < 2 || cnt > 3 {
					board[i][j] = 2
				}
			}

		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == -1 {
				board[i][j] = 1
			} else if board[i][j] == 2 {
				board[i][j] = 0
			}
		}
	}
}

func numOfLiveCell(board [][]int, i, j int) int {
	res := 0
	if i > 0 && board[i-1][j] > 0 {
		res++
	}
	if j > 0 && board[i][j-1] > 0 {
		res++
	}
	if i < len(board)-1 && board[i+1][j] > 0 {
		res++
	}
	if j < len(board[i])-1 && board[i][j+1] > 0 {
		res++
	}
	if i > 0 && j > 0 && board[i-1][j-1] > 0 {
		res++
	}
	if i < len(board)-1 && j < len(board[i])-1 && board[i+1][j+1] > 0 {
		res++
	}
	if i > 0 && j < len(board[i])-1 && board[i-1][j+1] > 0 {
		res++
	}
	if i < len(board)-1 && j > 0 && board[i+1][j-1] > 0 {
		res++
	}
	return res
}

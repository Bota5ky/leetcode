package offer

// https://leetcode.cn/problems/ju-zhen-zhong-de-lu-jing-lcof/
func exist(board [][]byte, word string) bool {
	if word == "" {
		return false
	}
	passed := make(map[[2]int]bool)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if dfs(board, i, j, []byte(word), 0, passed) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i int, j int, word []byte, p int, passed map[[2]int]bool) bool {
	if p >= len(word) {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}
	if board[i][j] != word[p] || passed[[2]int{i, j}] == true {
		return false
	}
	passed[[2]int{i, j}] = true
	if dfs(board, i-1, j, word, p+1, passed) {
		return true
	}
	if dfs(board, i+1, j, word, p+1, passed) {
		return true
	}
	if dfs(board, i, j-1, word, p+1, passed) {
		return true
	}
	if dfs(board, i, j+1, word, p+1, passed) {
		return true
	}
	passed[[2]int{i, j}] = false
	return false
}

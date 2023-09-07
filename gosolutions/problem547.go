package gosolutions

// https://leetcode-cn.com/problems/friend-circles/
func findCircleNum(M [][]int) int {
	cnt := 0
	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M[i]); j++ {
			if M[i][j] == 1 {
				cnt++
				friend(M, i)
			}
		}
	}
	return cnt
}

func friend(M [][]int, i int) {
	for j := 0; j < len(M[i]); j++ {
		if M[i][j] == 1 {
			M[i][j] = 0
			friend(M, j)
		}
	}
}

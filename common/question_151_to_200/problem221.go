package question_151_to_200

// https://leetcode.cn/problems/maximal-square/
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	pre := make([]int, len(matrix[0]))
	max := 0
	for i := 0; i < len(matrix[0]); i++ {
		pre[i] = int(matrix[0][i] - '0')
		if max < pre[i] {
			max = pre[i]
		}
	}
	for i := 1; i < len(matrix); i++ {
		temp := make([]int, len(matrix[0]))
		for j := 0; j < len(matrix[i]); j++ {
			if j == 0 {
				temp[j] = int(matrix[i][0] - '0')
			} else if matrix[i][j] == '0' {
				temp[j] = 0
			} else {
				temp[j] = min3(pre[j-1], temp[j-1], pre[j]) + 1
			}
			if max < temp[j] {
				max = temp[j]
			}
		}
		copy(pre, temp)
	}
	return max * max
}

func min3(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

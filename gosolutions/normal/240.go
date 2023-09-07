package leetcode

//https://leetcode-cn.com/problems/search-a-2d-matrix-ii/
func searchMatrix2(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	l := len(matrix[0])
	for i := 0; i < len(matrix); i++ {
		if matrix[i][l-1] < target {
			continue
		}
		if matrix[i][0] > target {
			break
		}
		for j, k := 0, l-1; j <= k; {
			mid := (j + k) / 2
			if matrix[i][mid] == target {
				return true
			} else if matrix[i][mid] > target {
				k = mid - 1
			} else {
				j = mid + 1
			}
		}
	}
	return false
}

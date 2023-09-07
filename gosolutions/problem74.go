package gosolutions

// https://leetcode-cn.com/problems/search-a-2d-matrix/
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	l := len(matrix[0])
	i, j := 0, len(matrix)*l-1
	for i <= j {
		mid := (i + j) / 2
		if matrix[mid/l][mid%l] == target {
			return true
		} else if matrix[mid/l][mid%l] > target {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return false
}

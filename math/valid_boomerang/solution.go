package valid_boomerang

// 1037. 有效的回旋镖 https://leetcode.cn/problems/valid-boomerang/
func isBoomerang(points [][]int) bool {
	if points[0][0] == points[1][0] && points[0][1] == points[1][1] {
		return false
	}
	if points[0][0] == points[2][0] && points[0][1] == points[2][1] {
		return false
	}
	if points[1][0] == points[2][0] && points[1][1] == points[2][1] {
		return false
	}
	if (points[1][1]-points[0][1])*(points[1][0]-points[2][0]) == (points[1][1]-points[2][1])*(points[1][0]-points[0][0]) {
		return false
	}
	return true
}

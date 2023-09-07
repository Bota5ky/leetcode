package gosolutions

// https://leetcode-cn.com/problems/check-if-n-and-its-double-exist/
func checkIfExist(arr []int) bool {
	m := make(map[int]bool)
	for _, v := range arr {
		if v == 0 && m[v] {
			return true
		}
		m[v] = true
		if v != 0 && (v%2 == 0 && m[v/2] || m[2*v]) {
			return true
		}
	}
	return false
}

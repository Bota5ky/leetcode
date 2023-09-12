package contains_duplicate

// https://leetcode.cn/problems/contains-duplicate/
func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, v := range nums {
		if m[v] {
			return true
		}
		m[v] = true
	}
	return false
}

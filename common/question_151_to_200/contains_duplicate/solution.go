package contains_duplicate

// 217. 存在重复元素 https://leetcode.cn/problems/contains-duplicate/
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

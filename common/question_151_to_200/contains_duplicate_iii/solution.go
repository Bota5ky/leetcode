package contains_duplicate_iii

// 220. 存在重复元素 III https://leetcode.cn/problems/contains-duplicate-iii/
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums) && j-i <= k; j++ {
			if abs(nums[i]-nums[j]) <= t {
				return true
			}
		}
	}
	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

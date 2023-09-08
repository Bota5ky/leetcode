package common

// https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/
func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	if len(nums2) < len(nums1) {
		nums1, nums2 = nums2, nums1
	}
	for i := 0; i < len(nums2); i++ {
		m[nums2[i]]++
	}
	for i := 0; i < len(nums1); i++ {
		if m[nums1[i]] == 0 {
			nums1 = append(nums1[:i], nums1[i+1:]...)
			i--
		}
	}
	return nums1
}

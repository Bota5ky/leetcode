package merge_sorted_array

// 88. 合并两个有序数组 https://leetcode.cn/problems/merge-sorted-array/
func merge2(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	for k := len(nums1) - 1; k >= 0; k-- {
		if j < 0 || k == i {
			break
		}
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
}

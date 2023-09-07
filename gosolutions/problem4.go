package gosolutions

// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/submissions/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	L := len(nums1) + len(nums2)
	return float64(findK(nums1, nums2, 0, 0, (L+1)/2)+findK(nums1, nums2, 0, 0, (L+2)/2)) / 2
}
func findK(nums1, nums2 []int, i, j, k int) int {
	if i >= len(nums1) {
		return nums2[j+k-1]
	}
	if j >= len(nums2) {
		return nums1[i+k-1]
	}
	if k == 1 {
		return min(nums1[i], nums2[j])
	}
	p1, p2 := i+k/2-1, j+k/2-1
	if p1 >= len(nums1) {
		p1 = len(nums1) - 1
	}
	if p2 >= len(nums2) {
		p2 = len(nums2) - 1
	}
	if nums1[p1] > nums2[p2] {
		return findK(nums1, nums2, i, j+k/2, k-(p2-j+1))
	}
	return findK(nums1, nums2, i+k/2, j, k-(p1-i+1))
}

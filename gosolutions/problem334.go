package gosolutions

// https://leetcode-cn.com/problems/increasing-triplet-subsequence/
func increasingTriplet(nums []int) bool {
	mid, small := 1<<31-1, 1<<31-1
	for i := 0; i < len(nums); i++ {
		if nums[i] <= small {
			small = nums[i]
		} else if nums[i] <= mid {
			mid = nums[i]
		} else {
			return true
		}
	}
	return false
}

package question_751_to_1000

// https://leetcode.cn/problems/sort-an-array/
func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	//quickSort
	i, j := 0, len(nums)-1
	pivot := nums[0]
	for i < j {
		for ; i < j && nums[j] >= pivot; j-- {
		}
		nums[i] = nums[j]
		for ; i < j && nums[i] <= pivot; i++ {
		}
		nums[j] = nums[i]
	}
	nums[i] = pivot
	sortArray(nums[:i])
	sortArray(nums[i+1:])
	return nums
}

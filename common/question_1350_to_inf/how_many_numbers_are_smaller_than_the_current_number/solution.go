package how_many_numbers_are_smaller_than_the_current_number

// https://leetcode.cn/problems/how-many-numbers-are-smaller-than-the-current-number/
func smallerNumbersThanCurrent(nums []int) []int {
	ret := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		cnt := 0
		for j := 0; j < len(nums); j++ {
			if nums[i] > nums[j] {
				cnt++
			}
		}
		ret[i] = cnt
	}
	return ret
}

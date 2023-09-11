package common

// https://leetcode.cn/problems/single-number-ii/
func singleNumber3(nums []int) int {
	once, twice := 0, 0
	for _, v := range nums {
		once = ^twice & (once ^ v)
		twice = ^once & (twice ^ v)
	}
	return once
}

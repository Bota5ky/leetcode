package single_number_ii

// 137. 只出现一次的数字 II https://leetcode.cn/problems/single-number-ii/
func singleNumber(nums []int) int {
	once, twice := 0, 0
	for _, v := range nums {
		once = ^twice & (once ^ v)
		twice = ^once & (twice ^ v)
	}
	return once
}

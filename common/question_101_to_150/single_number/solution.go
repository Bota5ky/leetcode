package single_number

// https://leetcode.cn/problems/single-number/
func singleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}

package single_number

// 136. 只出现一次的数字 https://leetcode.cn/problems/single-number/
func singleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}

package c_yi_huo_by_eric_345

// https://leetcode.cn/problems/swap-numbers-lcci/solution/c-yi-huo-by-eric-345/
func swapNumbers(numbers []int) []int {
	return func(a, b int) []int {
		return []int{b, a}
	}(numbers[0], numbers[1])
}

// a ^ b ^ b = a
// a ^ b ^ a = b
// a ^ a = 0
// 0 ^ a = a

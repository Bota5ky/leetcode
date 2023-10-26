package c_yi_huo_by_eric_345

// 面试题 16.01. 交换数字 C++ 异或 https://leetcode.cn/problems/swap-numbers-lcci/solution/c-yi-huo-by-eric-345/
func swapNumbers(n []int) []int {
	n[0] = n[0] ^ n[1]
	n[1] = n[0] ^ n[1]
	n[0] = n[0] ^ n[1]
	return n
}

// a ^ b ^ b = a
// a ^ b ^ a = b
// a ^ a = 0
// 0 ^ a = a

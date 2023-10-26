package sum_of_two_integers

// 371. 两整数之和 https://leetcode.cn/problems/sum-of-two-integers/
func getSum(a int, b int) int {
	for b != 0 {
		c := (a & b) << 1
		a = a ^ b
		b = c
	}
	return a
}

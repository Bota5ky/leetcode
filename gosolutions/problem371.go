package gosolutions

// https://leetcode-cn.com/problems/sum-of-two-integers/
func getSum(a int, b int) int {
	for b != 0 {
		c := (a & b) << 1 //进位
		a = a ^ b         //保留了不同的位
		b = c             //保留了相同的进位
	}
	return a
}

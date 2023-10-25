package climbing_stairs

// 70. 爬楼梯 https://leetcode.cn/problems/climbing-stairs/
func climbStairs(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	a, b, c := 1, 1, 0
	for n > 1 {
		c = a + b
		a = b
		b = c
		n--
	}
	return c
}

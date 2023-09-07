package leetcode

//https://leetcode-cn.com/problems/climbing-stairs/
func climbStairs(n int) int {
	//f[1]=1 f[2]=2
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

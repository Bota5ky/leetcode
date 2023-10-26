package three_steps_problem_lcci

// 面试题 08.01. 三步问题 https://leetcode.cn/problems/three-steps-problem-lcci/
func waysToStep(n int) int {
	a, b, c := 1, 2, 4
	if n <= 1 {
		return 1
	} else if n == 2 {
		return 2
	} else if n == 3 {
		return 4
	}
	for ; n > 3; n-- {
		temp := a + b + c
		a = b
		b = c
		c = temp % 1000000007
	}
	return c
}

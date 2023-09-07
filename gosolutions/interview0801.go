package gosolutions

// https://leetcode-cn.com/problems/three-steps-problem-lcci/
func waysToStep(n int) int {
	//f[1]=1;f[2]=2;f[3]=4;f[4]=7;f[5]=13
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

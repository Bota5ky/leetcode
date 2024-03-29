package fei_bo_na_qi_shu_lie_lcof

// LCR 126. 斐波那契数 https://leetcode.cn/problems/fei-bo-na-qi-shu-lie-lcof/
func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	var fn int
	fn1 := 1
	fn2 := 0
	for i := 2; i <= n; i++ {
		fn = fn1 + fn2
		fn2 = fn1
		fn1 = fn
	}
	return fn
}

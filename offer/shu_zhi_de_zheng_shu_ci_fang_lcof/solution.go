package offer

// 剑指 Offer 16. 数值的整数次方 https://leetcode.cn/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/
func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1.0 / x
		n = -n
	}
	return calc(x, n)
}

func calc(x float64, n int) float64 {
	if n == 1 {
		return x
	} else if n == 2 {
		return x * x
	}
	if n%2 == 0 {
		return calc(calc(x, n/2), 2)
	}
	return calc(calc(x, (n-1)/2), 2) * x
}

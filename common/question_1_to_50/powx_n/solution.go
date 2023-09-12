package powx_n

// 50. Pow(x, n) https://leetcode.cn/problems/powx-n/
// 剑指 Offer 16. 数值的整数次方 https://leetcode.cn/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/
func myPow2(x float64, n int) float64 {
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
	return calc2(x, n)
}

func calc2(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	}
	res := calc2(x, n>>1)
	if n&1 == 1 {
		return res * res * x
	}
	return res * res
}

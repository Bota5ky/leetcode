package offer

// 剑指 Offer 14- II. 剪绳子 II https://leetcode.cn/problems/jian-sheng-zi-ii-lcof/
func cuttingRope2(n int) int {
	switch n {
	case 2:
		return 1
	case 3:
		return 2
	}
	mul := 1
	for n > 4 {
		n -= 3
		mul *= 3
		if mul > 1000000007 {
			mul %= 1000000007
		}
	}
	return mul * n % 1000000007
}

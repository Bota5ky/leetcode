package integer_break

// 343. 整数拆分 https://leetcode.cn/problems/integer-break/
// LCR 131. 砍竹子 I https://leetcode.cn/problems/jian-sheng-zi-lcof/
func cuttingRope(n int) int {
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
	}
	return mul * n
}

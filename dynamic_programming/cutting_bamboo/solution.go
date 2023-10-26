package jian_sheng_zi_ii_lcof

// LCR 132. 砍竹子 II https://leetcode.cn/problems/jian-sheng-zi-ii-lcof/
func cuttingBamboo(n int) int {
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

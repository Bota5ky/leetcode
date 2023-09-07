package gosolutions

// https://leetcode-cn.com/problems/exchange-lcci/
func exchangeBits(num int) int {
	for i := 0; ; i += 2 {
		if num>>i == 0 {
			break
		}
		if num>>i&1 != num>>(i+1)&1 {
			if num>>i&1 == 1 {
				num -= 1 << i
				num += 1 << (i + 1)
			} else {
				num += 1 << i
				num -= 1 << (i + 1)
			}
		}
	}
	return num
}

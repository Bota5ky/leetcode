package reverse_bits_lcci

// https://leetcode-cn.com/problems/reverse-bits-lcci/
func reverseBits(num int) int {
	max, n1, n2 := 0, 0, 0
	for {
		if num&1 == 1 {
			n1++
		} else {
			if 1+n1+n2 > max {
				max = 1 + n1 + n2
			}
			if num == 0 {
				break
			}
			n2 = n1
			n1 = 0
		}
		num >>= 1
	}
	return max
}

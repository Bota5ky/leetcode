package leetcode

//https://leetcode-cn.com/problems/insert-into-bits-lcci/
func insertBits(N int, M int, i int, j int) int {
	for i > 0 {
		M <<= 1
		M = M | ((N >> (i - 1)) & 1)
		i--
	}
	N >>= (j + 1)
	N <<= (j + 1)
	return N | M
}

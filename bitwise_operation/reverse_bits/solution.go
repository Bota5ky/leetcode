package reverse_bits

// 190. 颠倒二进制位 https://leetcode.cn/problems/reverse-bits/
func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		a := num & 1
		res = (res << 1) + a
		num >>= 1
	}
	return res
}

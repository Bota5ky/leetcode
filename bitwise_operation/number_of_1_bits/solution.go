package number_of_1_bits

// 191. 位1的个数 https://leetcode.cn/problems/number-of-1-bits/
func hammingWeight(num uint32) int {
	cnt := 0
	for num != 0 {
		cnt += int(num) & 1
		num >>= 1
	}
	return cnt
}

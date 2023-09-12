package prime_number_of_set_bits_in_binary_representation

// 762. 二进制表示中质数个计算置位 https://leetcode.cn/problems/prime-number-of-set-bits-in-binary-representation/
func countPrimeSetBits(L int, R int) int {
	ret := 0
	for i := L; i <= R; i++ {
		num := i
		cnt := 0
		for num > 0 {
			if num&1 == 1 {
				cnt++
			}
			num >>= 1
		}
		if cnt == 2 || cnt == 3 || cnt == 5 || cnt == 7 || cnt == 11 || cnt == 13 || cnt == 17 || cnt == 19 {
			ret++
		}
	}
	return ret
}

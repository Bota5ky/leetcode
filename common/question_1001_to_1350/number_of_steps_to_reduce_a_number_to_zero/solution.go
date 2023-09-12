package number_of_steps_to_reduce_a_number_to_zero

// 1342. 将数字变成 0 的操作次数 https://leetcode.cn/problems/number-of-steps-to-reduce-a-number-to-zero/
func numberOfSteps(num int) int {
	cnt := 0
	for num > 0 {
		cnt += num&1 + 1
		num >>= 1
	}
	return cnt - 1
}

package question_1001_to_1350

// https://leetcode.cn/problems/number-of-steps-to-reduce-a-number-to-zero/
func numberOfSteps(num int) int {
	cnt := 0
	for num > 0 {
		cnt += num&1 + 1
		num >>= 1
	}
	return cnt - 1
}

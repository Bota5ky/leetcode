package gosolutions

// https://leetcode-cn.com/problems/number-of-steps-to-reduce-a-number-to-zero/
func numberOfSteps(num int) int {
	cnt := 0
	for num > 0 {
		cnt += num&1 + 1
		num >>= 1
	}
	return cnt - 1
}

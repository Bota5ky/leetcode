package daily_temperatures

// https://leetcode.cn/problems/daily-temperatures/
func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	stack := []int{0}
	for i := 1; i < len(T); i++ {
		if T[i] <= T[stack[len(stack)-1]] {
			stack = append(stack, i)
			continue
		}
		for len(stack) > 0 && T[stack[len(stack)-1]] < T[i] {
			res[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}

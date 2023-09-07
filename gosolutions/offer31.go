package gosolutions

// https://leetcode-cn.com/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/
func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	j := 0
	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

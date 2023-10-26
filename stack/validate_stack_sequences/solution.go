package validate_stack_sequences

// 946. 验证栈序列 https://leetcode.cn/problems/validate-stack-sequences/
// LCR 148. 验证图书取出顺序 https://leetcode.cn/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/
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

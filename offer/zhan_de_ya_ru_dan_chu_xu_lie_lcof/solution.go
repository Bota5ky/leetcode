package offer

// 剑指 Offer 31. 栈的压入、弹出序列 https://leetcode.cn/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/
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

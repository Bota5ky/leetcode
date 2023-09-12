package min_stack_lcci

// MinStack 剑指 Offer 30. 包含min函数的栈 https://leetcode.cn/problems/bao-han-minhan-shu-de-zhan-lcof/
// 面试题 03.02. 栈的最小值 https://leetcode.cn/problems/min-stack-lcci/
// 155. 最小栈 https://leetcode.cn/problems/min-stack/
type MinStack struct {
	data    []int
	support []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (t *MinStack) Push(x int) {
	t.data = append(t.data, x)
	k := len(t.support) - 1
	if k < 0 || t.support[k] >= x {
		t.support = append(t.support, x)
	} else {
		t.support = append(t.support, t.support[k])
	}
}

func (t *MinStack) Pop() {
	k := len(t.data)
	if k > 0 {
		t.data = t.data[:k-1]
		t.support = t.support[:k-1]
	}
}

func (t *MinStack) Top() int {
	return t.data[len(t.data)-1]
}

func (t *MinStack) Min() int {
	return t.support[len(t.support)-1]
}

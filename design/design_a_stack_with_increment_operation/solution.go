package design_a_stack_with_increment_operation

// CustomStack 1381. 设计一个支持增量操作的栈 https://leetcode.cn/problems/design-a-stack-with-increment-operation/
type CustomStack struct {
	stack   []int
	support []int
	max     int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{max: maxSize}
}

func (t *CustomStack) Push(x int) {
	if len(t.stack) < t.max {
		t.stack = append(t.stack, x)
	}
}

func (t *CustomStack) Pop() int {
	if len(t.stack) == 0 {
		return -1
	}
	res := t.stack[len(t.stack)-1]
	t.stack = t.stack[:len(t.stack)-1]
	return res
}

func (t *CustomStack) Increment(k int, val int) {
	for len(t.stack) > 0 {
		temp := t.stack[len(t.stack)-1]
		t.stack = t.stack[:len(t.stack)-1]
		t.support = append(t.support, temp)
	}
	for i := 0; i < k && len(t.support) > 0; i++ {
		temp := t.support[len(t.support)-1] + val
		t.support = t.support[:len(t.support)-1]
		t.Push(temp)
	}
	for len(t.support) > 0 {
		temp := t.support[len(t.support)-1]
		t.support = t.support[:len(t.support)-1]
		t.stack = append(t.stack, temp)
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */

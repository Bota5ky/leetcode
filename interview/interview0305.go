package interview

// SortedStack 面试题 03.05. 栈排序 https://leetcode.cn/problems/sort-of-stacks-lcci/
type SortedStack struct {
	stack   []int
	support []int
}

func ConstructorSortedStack() SortedStack {
	return SortedStack{}
}

func (t *SortedStack) Push(val int) {
	for len(t.stack) > 0 && val > t.stack[len(t.stack)-1] {
		temp := t.stack[len(t.stack)-1]
		t.support = append(t.support, temp)
		t.stack = t.stack[:len(t.stack)-1]
	}
	t.stack = append(t.stack, val)
	for len(t.support) > 0 {
		temp := t.support[len(t.support)-1]
		t.stack = append(t.stack, temp)
		t.support = t.support[:len(t.support)-1]
	}
}

func (t *SortedStack) Pop() {
	if len(t.stack) == 0 {
		return
	}
	t.stack = t.stack[:len(t.stack)-1]
}

func (t *SortedStack) Peek() int {
	if len(t.stack) == 0 {
		return -1
	}
	return t.stack[len(t.stack)-1]
}

func (t *SortedStack) IsEmpty() bool {
	return len(t.stack) == 0
}

/**
 * Your SortedStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.IsEmpty();
 */

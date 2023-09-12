package implement_stack_using_queues

// MyStack 225. 用队列实现栈 https://leetcode.cn/problems/implement-stack-using-queues/
type MyStack struct {
	q1 []int
	q2 []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (t *MyStack) Push(x int) {
	t.q1 = append(t.q1, x)
}

func (t *MyStack) Pop() int {
	for len(t.q1) > 1 {
		t.q2 = append(t.q2, t.q1[0])
		t.q1 = t.q1[1:]
	}
	ret := t.q1[0]
	t.q1 = t.q2
	t.q2 = []int{}
	return ret
}

func (t *MyStack) Top() int {
	for len(t.q1) > 1 {
		t.q2 = append(t.q2, t.q1[0])
		t.q1 = t.q1[1:]
	}
	ret := t.q1[0]
	t.q2 = append(t.q2, ret)
	t.q1 = t.q2
	t.q2 = []int{}
	return ret
}

func (t *MyStack) Empty() bool {
	return len(t.q1) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

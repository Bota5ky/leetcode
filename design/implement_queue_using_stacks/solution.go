package implement_queue_using_stacks

// MyQueue 232. 用栈实现队列 https://leetcode.cn/problems/implement-queue-using-stacks/
// LCR 125. 图书整理 II https://leetcode.cn/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/
// 面试题 03.04. 化栈为队 https://leetcode.cn/problems/implement-queue-using-stacks-lcci/
type MyQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() MyQueue {
	return MyQueue{
		make([]int, 0),
		make([]int, 0),
	}
}

func (t *MyQueue) AppendTail(value int) {
	t.stack1 = append(t.stack1, value)
}

func (t *MyQueue) DeleteHead() int {
	if len(t.stack2) == 0 {
		for len(t.stack1) > 0 {
			t.stack2 = append(t.stack2, t.stack1[len(t.stack1)-1])
			t.stack1 = t.stack1[:len(t.stack1)-1]
		}
	}
	if len(t.stack2) == 0 {
		return -1
	}
	ret := t.stack2[len(t.stack2)-1]
	t.stack2 = t.stack2[:len(t.stack2)-1]
	return ret
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

package gosolutions

// CQueue CQueue
// https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/
type CQueue struct {
	stack1 []int
	stack2 []int
}

// Constructor Constructor
func Constructor() CQueue {
	return CQueue{
		make([]int, 0),
		make([]int, 0),
	}
}

// AppendTail AppendTail
func (t *CQueue) AppendTail(value int) {
	t.stack1 = append(t.stack1, value)
}

// DeleteHead DeleteHead
func (t *CQueue) DeleteHead() int {
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
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

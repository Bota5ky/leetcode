package yong_liang_ge_zhan_shi_xian_dui_lie_lcof

// CQueue 剑指 Offer 09. 用两个栈实现队列 https://leetcode.cn/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/
type CQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() CQueue {
	return CQueue{
		make([]int, 0),
		make([]int, 0),
	}
}

func (t *CQueue) AppendTail(value int) {
	t.stack1 = append(t.stack1, value)
}

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

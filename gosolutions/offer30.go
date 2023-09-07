package gosolutions

// MinStack MinStack
// https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/
type MinStack struct {
	data    []int
	support []int
}

// Constructor2 new
func Constructor2() MinStack {
	return MinStack{}
}

// Push Push
func (t *MinStack) Push(x int) {
	t.data = append(t.data, x)
	k := len(t.support) - 1
	if k < 0 || t.support[k] >= x {
		t.support = append(t.support, x)
	} else {
		t.support = append(t.support, t.support[k])
	}
}

// Pop Pop
func (t *MinStack) Pop() {
	k := len(t.data)
	if k > 0 {
		t.data = t.data[:k-1]
		t.support = t.support[:k-1]
	}
}

// Top Top
func (t *MinStack) Top() int {
	return t.data[len(t.data)-1]
}

// Min Min
func (t *MinStack) Min() int {
	return t.support[len(t.support)-1]
}

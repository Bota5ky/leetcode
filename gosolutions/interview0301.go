package gosolutions

// TripleInOne Triple []int
// https://leetcode-cn.com/problems/three-in-one-lcci/
type TripleInOne struct {
	stack []int
	//0+3i 1+3i 2+3i  0 1 2记录数目 3 4 5 开始存数
}

// Constructor5 Constructor
func Constructor5(stackSize int) TripleInOne {
	var res TripleInOne
	res.stack = make([]int, 3*stackSize+3)
	return res
}

// Push Push
func (t *TripleInOne) Push(stackNum int, value int) {
	if t.stack[stackNum] < (len(t.stack)-3)/3 {
		t.stack[stackNum]++
		t.stack[stackNum+3*t.stack[stackNum]] = value
	}
}

// Pop Pop
func (t *TripleInOne) Pop(stackNum int) int {
	if t.stack[stackNum] == 0 {
		return -1
	}
	ret := t.stack[stackNum+3*t.stack[stackNum]]
	t.stack[stackNum]--
	return ret
}

// Peek Peek
func (t *TripleInOne) Peek(stackNum int) int {
	if t.stack[stackNum] == 0 {
		return -1
	}
	return t.stack[stackNum+3*t.stack[stackNum]]
}

// IsEmpty IsEmpty
func (t *TripleInOne) IsEmpty(stackNum int) bool {
	return t.stack[stackNum] == 0
}

/**
 * Your TripleInOne object will be instantiated and called as such:
 * obj := Constructor(stackSize);
 * obj.Push(stackNum,value);
 * param_2 := obj.Pop(stackNum);
 * param_3 := obj.Peek(stackNum);
 * param_4 := obj.IsEmpty(stackNum);
 */

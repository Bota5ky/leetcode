package leetcode

//StackOfPlates StackOfPlates
//https://leetcode-cn.com/problems/stack-of-plates-lcci/
type StackOfPlates struct {
	size   int
	stacks [][]int
}

//Constructor6 Constructor6
func Constructor6(cap int) StackOfPlates {
	return StackOfPlates{size: cap}
}

//Push Push
func (t *StackOfPlates) Push(val int) {
	if t.size <= 0 {
		return
	}
	if len(t.stacks) == 0 || len(t.stacks[len(t.stacks)-1]) == t.size {
		//为空或者栈满了 新建
		t.stacks = append(t.stacks, []int{})
	}
	t.stacks[len(t.stacks)-1] = append(t.stacks[len(t.stacks)-1], val)
}

//Pop Pop
func (t *StackOfPlates) Pop() int {
	if len(t.stacks) == 0 || len(t.stacks[len(t.stacks)-1]) == 0 {
		return -1
	}
	res := t.stacks[len(t.stacks)-1][len(t.stacks[len(t.stacks)-1])-1]
	t.stacks[len(t.stacks)-1] = t.stacks[len(t.stacks)-1][:len(t.stacks[len(t.stacks)-1])-1]
	if len(t.stacks[len(t.stacks)-1]) == 0 {
		t.stacks = t.stacks[:len(t.stacks)-1]
	}
	return res
}

//PopAt PopAt
func (t *StackOfPlates) PopAt(index int) int {
	if index < 0 || index >= len(t.stacks) || len(t.stacks[index]) == 0 {
		return -1
	}
	res := t.stacks[index][len(t.stacks[index])-1]
	t.stacks[index] = t.stacks[index][:len(t.stacks[index])-1]
	if len(t.stacks[index]) == 0 {
		t.stacks = append(t.stacks[:index], t.stacks[index+1:]...)
	}
	return res
}

/**
 * Your StackOfPlates object will be instantiated and called as such:
 * obj := Constructor(cap);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAt(index);
 */

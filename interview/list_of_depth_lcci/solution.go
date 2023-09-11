package list_of_depth_lcci

import . "leetcode/model"

// https://leetcode-cn.com/problems/list-of-depth-lcci/
func listOfDepth(tree *TreeNode) []*ListNode {
	if tree == nil {
		return []*ListNode{}
	}
	stack := []*TreeNode{tree}
	var ret []*ListNode
	for len(stack) > 0 {
		head := &ListNode{Val: stack[0].Val}
		temp := head
		k := len(stack)
		for i := 1; i < k; i++ {
			head.Next = &ListNode{Val: stack[i].Val}
			head = head.Next
		}
		head.Next = nil
		ret = append(ret, temp)
		for i := 0; i < k; i++ {
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
		}
		stack = stack[k:]
	}
	return ret
}

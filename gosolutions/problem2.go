package gosolutions

// https://leetcode-cn.com/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	plus := 0
	pre := &ListNode{}
	ret := pre
	for l1 != nil || l2 != nil || plus != 0 {
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		sum := a + b + plus
		node := &ListNode{Val: sum % 10}
		plus = sum / 10
		pre.Next = node
		pre = node
	}
	return ret.Next
}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

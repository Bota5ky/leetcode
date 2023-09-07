package gosolutions

import . "gosolutions/model"

// 数字反向存放
// https://leetcode-cn.com/problems/sum-lists-lcci/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	plus := 0
	pre := &ListNode{}
	ret := pre
	for plus != 0 || l1 != nil || l2 != nil {
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		pre.Next = &ListNode{Val: (a + b + plus) % 10, Next: nil}
		plus = (a + b + plus) / 10
		pre = pre.Next
	}
	return ret.Next
}

//数字正向存放

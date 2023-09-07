package gosolutions

import . "gosolutions/model"

// 面试题 02.05. 链表求和 https://leetcode.cn/problems/sum-lists-lcci/
// 2. 两数相加 https://leetcode.cn/problems/add-two-numbers/
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

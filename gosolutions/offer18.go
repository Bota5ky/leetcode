package gosolutions

import . "gosolutions/model"

// https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/
func deleteNodeByValue(head *ListNode, val int) *ListNode {
	mark := &ListNode{0, head}
	res := mark
	for head != nil {
		if head.Val == val {
			mark.Next = mark.Next.Next
			break
		}
		mark = head
		head = head.Next
	}
	return res.Next
}

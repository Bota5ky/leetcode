package offer

import . "leetcode/model"

// 剑指 Offer 24. 反转链表 https://leetcode.cn/problems/fan-zhuan-lian-biao-lcof/
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prehead *ListNode
	rear := head.Next
	for head != nil {
		head.Next = prehead
		prehead = head
		head = rear
		if rear != nil {
			rear = rear.Next
		}
	}
	return prehead
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

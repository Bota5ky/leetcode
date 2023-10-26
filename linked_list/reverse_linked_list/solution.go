package reverse_linked_list

import . "leetcode/_model"

// 206. 反转链表 https://leetcode.cn/problems/reverse-linked-list/
// LCR 141. 训练计划 III https://leetcode.cn/problems/fan-zhuan-lian-biao-lcof/
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var preHead *ListNode
	rear := head.Next
	for head != nil {
		head.Next = preHead
		preHead = head
		head = rear
		if rear != nil {
			rear = rear.Next
		}
	}
	return preHead
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

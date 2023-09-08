package interview

import . "gosolutions/model"

// 面试题 02.04. 分割链表 https://leetcode.cn/problems/partition-list-lcci/
func partition(head *ListNode, x int) *ListNode {
	res := head
	rear := head
	for head != nil && rear != nil {
		for head != nil && head.Val < x {
			head = head.Next
		}
		if head != nil {
			rear = head.Next
		}
		for rear != nil && rear.Val >= x {
			rear = rear.Next
		}
		if rear != nil && head != nil {
			head.Val, rear.Val = rear.Val, head.Val
		}
	}
	return res
}

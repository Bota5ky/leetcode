package gosolutions

import . "gosolutions/model"

// https://leetcode-cn.com/problems/partition-list-lcci/
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

package gosolutions

import . "gosolutions/model"

// https://leetcode-cn.com/problems/linked-list-cycle-lcci/
func detectCycle(head *ListNode) *ListNode {
	slow, fast, rec := head, head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	fast = rec
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

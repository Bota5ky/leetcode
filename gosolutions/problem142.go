package gosolutions

// https://leetcode-cn.com/problems/linked-list-cycle-ii/
func detectCycle(head *ListNode) *ListNode {
	rem, slow, fast := head, head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	slow = rem
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

package linked_list_cycle

// https://leetcode.cn/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for {
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	return true
}

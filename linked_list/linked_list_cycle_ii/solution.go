package linked_list_cycle_ii

import . "leetcode/_model"

// 142. 环形链表 II https://leetcode.cn/problems/linked-list-cycle-ii/
// 面试题 02.08. 环路检测 https://leetcode.cn/problems/linked-list-cycle-lcci/
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

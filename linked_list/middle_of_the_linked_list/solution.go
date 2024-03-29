package middle_of_the_linked_list

import . "leetcode/_model"

// 876. 链表的中间结点 https://leetcode.cn/problems/middle-of-the-linked-list/
func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

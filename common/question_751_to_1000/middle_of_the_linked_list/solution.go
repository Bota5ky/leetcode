package middle_of_the_linked_list

// https://leetcode.cn/problems/middle-of-the-linked-list/
func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

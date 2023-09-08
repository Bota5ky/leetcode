package common

// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
func reverseKGroup(head *ListNode, k int) *ListNode {
	temp := head
	for i := k; i > 0; i-- {
		if temp == nil {
			return head
		}
		temp = temp.Next
	}
	pre := reverseKGroup(temp, k)
	rear := head.Next
	for i := k; i > 0; i-- {
		head.Next = pre
		pre = head
		head = rear
		if rear != nil {
			rear = rear.Next
		}
	}
	return pre
}

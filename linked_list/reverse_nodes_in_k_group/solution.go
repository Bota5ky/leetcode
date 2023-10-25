package reverse_nodes_in_k_group

import . "leetcode/model"

// 25. K 个一组翻转链表 https://leetcode.cn/problems/reverse-nodes-in-k-group/
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

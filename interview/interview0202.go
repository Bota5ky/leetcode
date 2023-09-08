package interview

import . "gosolutions/model"

// 面试题 02.02. 返回倒数第 k 个节点 https://leetcode.cn/problems/kth-node-from-end-of-list-lcci/
func kthToLast(head *ListNode, k int) int {
	rear := head
	for rear != nil && k > 0 {
		rear = rear.Next
		k--
	}
	for rear != nil {
		head = head.Next
		rear = rear.Next
	}
	return head.Val
}

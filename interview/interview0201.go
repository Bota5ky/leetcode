package interview

import . "gosolutions/model"

// 面试题 02.01. 移除重复节点 https://leetcode.cn/problems/remove-duplicate-node-lcci/
func removeDuplicateNodes(head *ListNode) *ListNode {
	ret := head
	for head != nil {
		pre := head
		temp := head.Val
		for pre.Next != nil {
			if pre.Next.Val == temp {
				pre.Next = pre.Next.Next
			} else {
				pre = pre.Next
			}
		}
		head = head.Next
	}
	return ret
}

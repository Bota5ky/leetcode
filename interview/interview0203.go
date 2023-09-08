package interview

import . "gosolutions/model"

// 面试题 02.03. 删除中间节点 https://leetcode.cn/problems/delete-middle-node-lcci/
// 237. 删除链表中的节点 https://leetcode.cn/problems/delete-node-in-a-linked-list/
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

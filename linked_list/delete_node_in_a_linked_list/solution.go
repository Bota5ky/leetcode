package delete_node_in_alinked_list

import . "leetcode/_model"

// 237. 删除链表中的节点 https://leetcode.cn/problems/delete-node-in-a-linked-list/
// 面试题 02.03. 删除中间节点 https://leetcode.cn/problems/delete-middle-node-lcci/
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

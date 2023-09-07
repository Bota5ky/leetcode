package leetcode

//https://leetcode-cn.com/problems/delete-middle-node-lcci/
func deleteNode2(node *ListNode) {
	*node = *node.Next
}

// func deleteNode(node *ListNode) {
//     node.Val=node.Next.Val
//     node.Next=node.Next.Next
// }

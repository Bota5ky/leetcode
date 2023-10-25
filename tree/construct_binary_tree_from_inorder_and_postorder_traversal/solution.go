package construct_binary_tree_from_inorder_and_postorder_traversal

import . "leetcode/model"

// 106. 从中序与后序遍历序列构造二叉树 https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	if len(inorder) == 1 {
		node := &TreeNode{Val: inorder[0]}
		return node
	}
	last := postorder[len(postorder)-1]
	root := &TreeNode{Val: last}
	i := 0
	for ; i < len(inorder) && inorder[i] != last; i++ {
	}
	root.Left = buildTree(inorder[:i], postorder[:i])
	root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
	return root
}

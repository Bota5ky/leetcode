package common

// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
func buildTree2(inorder []int, postorder []int) *TreeNode {
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

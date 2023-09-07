package leetcode

//TreeNode Definition for a binary tree node.
//https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//返回值
//每次作甚
//退出条件
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	mid := preorder[0]
	var head TreeNode
	head.Val = mid
	var i int
	for i = 0; inorder[i] != mid; i++ {
	}
	m := len(inorder[:i])
	n := len(inorder[i+1:])
	head.Left = buildTree(preorder[1:m+1], inorder[:i])
	head.Right = buildTree(preorder[m+1:m+n+1], inorder[i+1:])
	return &head
}

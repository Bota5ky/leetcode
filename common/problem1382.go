package common

// https://leetcode-cn.com/problems/balance-a-binary-search-tree/
func balanceBST(root *TreeNode) *TreeNode {
	var ret []int
	list(root, &ret)
	return createTree(ret)
}

// 中序遍历
func list(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}
	list(root.Left, ret)
	*ret = append(*ret, root.Val)
	list(root.Right, ret)
}

func createTree(ret []int) *TreeNode {
	if len(ret) == 0 {
		return nil
	}
	mid := len(ret) / 2
	node := &TreeNode{Val: ret[mid]}
	node.Left = createTree(ret[:mid])
	node.Right = createTree(ret[mid+1:])
	return node
}

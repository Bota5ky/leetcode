package binary_tree_paths

import (
	. "leetcode/model"
	"strconv"
)

// 257. 二叉树的所有路径 https://leetcode.cn/problems/binary-tree-paths/
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	var res []string
	temp := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		res = append(res, temp)
		return res
	}
	routine(root.Left, temp, &res)
	routine(root.Right, temp, &res)
	return res
}

// 之前积累的字符串
func routine(node *TreeNode, temp string, res *[]string) {
	if node == nil {
		return
	}
	num := strconv.Itoa(node.Val)
	temp += "->" + num
	if node.Left == nil && node.Right == nil {
		*res = append(*res, temp)
		return
	}
	routine(node.Left, temp, res)
	routine(node.Right, temp, res)
}

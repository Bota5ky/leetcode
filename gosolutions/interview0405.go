package gosolutions

// 中序遍历 是否严格升序
// https://leetcode-cn.com/problems/legal-binary-search-tree-lcci/
func isValidBST(root *TreeNode) bool {
	var vals []int
	midScan(root, &vals)
	for i := 1; i < len(vals); i++ {
		if vals[i] <= vals[i-1] {
			return false
		}
	}
	return true
}

func midScan(root *TreeNode, vals *[]int) {
	if root == nil {
		return
	}
	midScan(root.Left, vals)
	*vals = append(*vals, root.Val)
	midScan(root.Right, vals)
}

// 不用格外数组的中序遍历
// var pre int

// func isValidBST(root *TreeNode) bool {
//     pre = -2 << 31
//     return ldr(root)
// }

// func ldr(root *TreeNode) bool {
//     if root == nil {
//         return true
//     }
//     if !ldr(root.Left) {
//         return false
//     }
//     if root.Val <= pre {
//         return false
//     }
//     pre = root.Val
//     return ldr(root.Right)
// }

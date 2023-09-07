package gosolutions

// https://leetcode-cn.com/problems/count-complete-tree-nodes/
func countNodes(root *TreeNode) int {
	stack := []*TreeNode{}
	sum := 0
	for root != nil || len(stack) > 0 {
		if root != nil {
			sum++
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1].Right
			stack = stack[:len(stack)-1]
		}
	}
	return sum
}

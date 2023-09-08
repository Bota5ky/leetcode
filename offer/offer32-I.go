package offer

// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
func levelOrder(root *TreeNode) []int {
	var ret []int
	j := 0
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		k := len(stack)
		for i := 0; i < k; i++ {
			if stack[i] == nil {
				continue
			}
			ret = append(ret, stack[i].Val)
			stack = append(stack, stack[i].Left, stack[i].Right)
		}
		stack = stack[k:]
		j++
	}
	return ret
}

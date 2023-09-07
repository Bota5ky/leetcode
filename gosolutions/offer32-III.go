package gosolutions

// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/
func levelOrder3(root *TreeNode) [][]int {
	var ret [][]int
	stack := []*TreeNode{root}
	j := 0
	for len(stack) != 0 {
		k := len(stack)
		flag := 0
		for i := 0; i < k; i++ {
			if stack[i] != nil {
				if flag == 0 {
					ret = append(ret, []int{})
					flag = 1
				}
				if j%2 == 0 {
					ret[j] = append(ret[j], stack[i].Val)
				} else {
					ret[j] = append([]int{stack[i].Val}, ret[j]...)
				}
				stack = append(stack, stack[i].Left, stack[i].Right)
			}
		}
		stack = stack[k:]
		j++
	}
	return ret
}

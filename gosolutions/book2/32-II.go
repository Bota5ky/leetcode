package leetcode

//https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/
func levelOrder2(root *TreeNode) [][]int {
	var ret [][]int
	j := 0
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		k := len(stack)
		flag:=0
		for i := 0; i < k; i++ {
			if stack[i]==nil {continue}
			if flag==0 {
				ret=append(ret,[]int{})
				flag=1
			}
			ret[j] = append(ret[j], stack[i].Val)
			stack = append(stack, stack[i].Left, stack[i].Right)
		}
		stack = stack[k:]
		j++
	}
	return ret
}
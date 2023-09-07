package leetcode

//https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/
func verifyPostorder(postorder []int) bool {
	j := len(postorder) - 1
	if j < 0 {
		return true
	}
	p := j
	for k := 0; k < len(postorder); k++ {
		if postorder[k] > postorder[j] {
			p = k
			break
		}
	}
	for k := p + 1; k < len(postorder); k++ {
		if postorder[k] < postorder[j] {
			return false
		}
	}
	return verifyPostorder(postorder[:p]) && verifyPostorder(postorder[p:j])
}

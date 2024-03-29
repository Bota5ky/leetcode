package er_cha_sou_suo_shu_de_hou_xu_bian_li_xu_lie_lcof

// LCR 152. 验证二叉搜索树的后序遍历序列 https://leetcode.cn/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/
func verifyTreeOrder(postorder []int) bool {
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
	return verifyTreeOrder(postorder[:p]) && verifyTreeOrder(postorder[p:j])
}

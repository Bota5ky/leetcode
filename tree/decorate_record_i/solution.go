package cong_shang_dao_xia_da_yin_er_cha_shu_lcof

import . "leetcode/_model"

// LCR 149. 彩灯装饰记录 I https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
func decorateRecord(root *TreeNode) []int {
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

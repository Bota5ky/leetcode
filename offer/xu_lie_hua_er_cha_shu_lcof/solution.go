package xu_lie_hua_er_cha_shu_lcof

import (
	. "leetcode/model"
	"strconv"
	"strings"
)

// Codec 剑指 Offer 37. 序列化二叉树 https://leetcode.cn/problems/xu-lie-hua-er-cha-shu-lcof/  TODO
func Codec(root *TreeNode) string {
	var ret []string
	stack := []*TreeNode{root}
	next := 0
	if root != nil {
		next++
	}
	for next > 0 {
		k := len(stack)
		next = 0
		for i := 0; i < k; i++ {
			if stack[i] != nil {
				stack = append(stack, stack[i].Left, stack[i].Right)
				ret = append(ret, strconv.Itoa(stack[i].Val))
				if stack[i].Left != nil {
					next++
				}
				if stack[i].Right != nil {
					next++
				}
			} else {
				stack = append(stack, nil, nil)
				ret = append(ret, "null")
			}
		}
		stack = stack[k:]
	}
	return "[" + strings.Join(ret, ",") + "]"
}

package cong_wei_dao_tou_da_yin_lian_biao_lcof

import . "leetcode/_model"

// LCR 123. 图书整理 I https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
func reverseBookList(head *ListNode) []int {
	var ret []int
	for head != nil {
		ret = append([]int{head.Val}, ret...)
		head = head.Next
	}
	return ret
}

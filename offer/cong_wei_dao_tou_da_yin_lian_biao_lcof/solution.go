package cong_wei_dao_tou_da_yin_lian_biao_lcof

import . "leetcode/model"

// 剑指 Offer 06. 从尾到头打印链表 https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
func reversePrint(head *ListNode) []int {
	var ret []int
	for head != nil {
		ret = append([]int{head.Val}, ret...)
		head = head.Next
	}
	return ret
}

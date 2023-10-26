package shan_chu_lian_biao_de_jie_dian_lcof

import . "leetcode/_model"

// LCR 136. 删除链表的节点 https://leetcode.cn/problems/shan-chu-lian-biao-de-jie-dian-lcof/
func deleteNode(head *ListNode, val int) *ListNode {
	mark := &ListNode{Next: head}
	res := mark
	for head != nil {
		if head.Val == val {
			mark.Next = mark.Next.Next
			break
		}
		mark = head
		head = head.Next
	}
	return res.Next
}

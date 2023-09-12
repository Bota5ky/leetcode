package liang_ge_lian_biao_de_di_yi_ge_gong_gong_jie_dian_lcof

import . "leetcode/model"

// 剑指 Offer 52. 两个链表的第一个公共节点 https://leetcode.cn/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/
// 160. 相交链表 https://leetcode.cn/problems/intersection-of-two-linked-lists/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a := headA
	b := headB
	for a != nil || b != nil {
		if a == nil {
			a = headB
			b = b.Next
			continue
		}
		if b == nil {
			b = headA
			a = a.Next
			continue
		}
		if a == b {
			return a
		}
		a = a.Next
		b = b.Next
	}
	return nil
}

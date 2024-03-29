package he_bing_liang_ge_pai_xu_de_lian_biao_lcof

import . "leetcode/_model"

// 21. 合并两个有序链表 https://leetcode.cn/problems/merge-two-sorted-lists/
// LCR 142. 训练计划 IV https://leetcode.cn/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := &ListNode{}
	node := ret
	for l1 != nil || l2 != nil {
		if l2 == nil || (l1 != nil && l1.Val < l2.Val) {
			node.Next = l1
			l1 = l1.Next
			node = node.Next
		} else {
			node.Next = l2
			l2 = l2.Next
			node = node.Next
		}
	}
	return ret.Next
}

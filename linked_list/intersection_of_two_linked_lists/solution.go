package intersection_of_two_linked_lists

import . "leetcode/_model"

// 160. 相交链表 https://leetcode.cn/problems/intersection-of-two-linked-lists/
// LCR 171. 训练计划 V https://leetcode.cn/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	recA, recB := headA, headB
	for headA != nil || headB != nil {
		if headA == headB {
			return headA
		}
		if headA == nil {
			headA = recB
		} else {
			headA = headA.Next
		}
		if headB == nil {
			headB = recA
		} else {
			headB = headB.Next
		}
	}
	return nil
}

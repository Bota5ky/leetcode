package question_151_to_200

// 和offer52相同
// https://leetcode.cn/problems/intersection-of-two-linked-lists/
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

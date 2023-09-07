package leetcode

//https://leetcode-cn.com/problems/intersection-of-two-linked-lists-lcci/
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	a := headA
	b := headB
	for headA != nil || headB != nil {
		if headA == headB {
			break
		}
		if headA == nil {
			headA = b
		} else {
			headA = headA.Next
		}
		if headB == nil {
			headB = a
		} else {
			headB = headB.Next
		}
	}
	return headA
}

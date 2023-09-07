package leetcode

//https://leetcode-cn.com/problems/insertion-sort-list/
func insertionSortList(head *ListNode) *ListNode {
	rem := &ListNode{}
	for head != nil {
		temp := head.Next
		pre := rem
		rear := pre.Next
		for rear != nil && rear.Val < head.Val {
			pre = pre.Next
			rear = rear.Next
		}
		head.Next = rear
		pre.Next = head
		head = temp
	}
	return rem.Next
}

package gosolutions

// https://leetcode-cn.com/problems/odd-even-linked-list/
func oddEvenList(head *ListNode) *ListNode {
	odd := &ListNode{}
	ret := odd
	even := &ListNode{}
	conn := even
	for i := 1; head != nil; i++ {
		next := head.Next
		if i%2 == 1 {
			odd.Next = head
			odd = odd.Next
			head.Next = nil
		} else {
			even.Next = head
			even = even.Next
			head.Next = nil
		}
		head = next
	}
	odd.Next = conn.Next
	return ret.Next
}

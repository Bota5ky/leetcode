package gosolutions

// https://leetcode-cn.com/problems/rotate-list/
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	rec, length := head, 0
	for {
		length++
		if head.Next == nil {
			head.Next = rec
			break
		}
		head = head.Next
	}
	k %= length
	k = length - k - 1
	for ; k > 0; k-- {
		rec = rec.Next
	}
	ret := rec.Next
	rec.Next = nil
	return ret
}

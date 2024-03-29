package insertion_sort_list

import . "leetcode/_model"

// 147. 对链表进行插入排序 https://leetcode.cn/problems/insertion-sort-list/
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

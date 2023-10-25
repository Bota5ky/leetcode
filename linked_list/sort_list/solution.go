package sort_list

import . "leetcode/model"

// 148. 排序链表 https://leetcode.cn/problems/sort-list/
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := &ListNode{Next: head}
	temp, slow, fast := head, head, head
	for fast != nil && fast.Next != nil {
		pre = pre.Next
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = nil
	return mergeList(sortList(temp), sortList(slow))
}

func mergeList(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{}
	ret := pre
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}
	if l2 != nil {
		pre.Next = l2
	}
	if l1 != nil {
		pre.Next = l1
	}
	return ret.Next
}

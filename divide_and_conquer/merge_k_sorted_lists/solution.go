package merge_k_sorted_lists

import . "leetcode/_model"

// 23. 合并 K 个升序链表 https://leetcode.cn/problems/merge-k-sorted-lists/
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for len(lists) > 1 {
		temp := make([]*ListNode, 0)
		for i := 0; i < len(lists); i += 2 {
			if i == len(lists)-1 {
				temp = append(temp, lists[i])
			} else {
				temp = append(temp, mergeTwoLists(lists[i], lists[i+1]))
			}
		}
		lists = temp
	}
	return lists[0]
}

func mergeTwoLists(a, b *ListNode) *ListNode {
	pre := &ListNode{}
	ret := pre
	for a != nil && b != nil {
		if a.Val < b.Val {
			pre.Next = a
			a = a.Next
		} else {
			pre.Next = b
			b = b.Next
		}
		pre = pre.Next
	}
	if a != nil {
		pre.Next = a
	}
	if b != nil {
		pre.Next = b
	}
	return ret.Next
}

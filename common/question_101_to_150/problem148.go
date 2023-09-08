package common

// https://leetcode-cn.com/problems/sort-list/
// 递归归并 时间复杂度 nlogn 空间复杂度 nlogn
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

// 迭代归并 时间复杂度 nlogn 空间复杂度 O(1)
func sortList2(head *ListNode) *ListNode {
	pre, n := &ListNode{Next: head}, 0
	for head != nil {
		n++ //链表总长度
		head = head.Next
	}
	for i := 1; i < n; i *= 2 {
		prehead := pre
		for prehead.Next != nil {
			l1 := prehead.Next
			cur := l1
			for j := 0; j < i-1 && cur != nil; j++ {
				cur = cur.Next
			}
			if cur == nil { // l2没有就没必要merge排序，直接退出
				break
			}
			l2 := cur.Next
			cur.Next = nil
			cur = l2
			for j := 0; j < i-1 && cur != nil; j++ {
				cur = cur.Next
			}
			var l3 *ListNode // l3赋值
			if cur != nil {
				l3 = cur.Next
				cur.Next = nil
			}
			node := mergeList(l1, l2)
			prehead.Next = node // pre 接上排序后的头节点
			for node.Next != nil {
				node = node.Next
			}
			node.Next = l3 // 接上排序后的尾节点
			prehead = node // pre 变为排序后的尾节点
		}
	}
	return pre.Next
}

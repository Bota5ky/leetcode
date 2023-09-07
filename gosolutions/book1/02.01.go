package leetcode

//不使用额外空间的暴力方法 O(N^2)
//https://leetcode-cn.com/problems/remove-duplicate-node-lcci/
func removeDuplicateNodes(head *ListNode) *ListNode {
	ret := head
	for head != nil {
		pre := head
		temp := head.Val
		for pre.Next != nil {
			if pre.Next.Val == temp {
				pre.Next = pre.Next.Next
			} else {
				pre = pre.Next
			}
		}
		head = head.Next
	}
	return ret
}

//ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

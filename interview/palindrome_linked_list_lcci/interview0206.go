package palindrome_linked_list_lcci

import . "leetcode/model"

// 面试题 02.06. 回文链表 https://leetcode.cn/problems/palindrome-linked-list-lcci/
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	//快慢指针
	l1 := head
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	l2 := reverseNode(slow)
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}

func reverseNode(head *ListNode) *ListNode {
	var pre, rear *ListNode
	if head != nil {
		rear = head.Next
	}
	for head != nil {
		head.Next = pre
		pre = head
		head = rear
		if rear != nil {
			rear = rear.Next
		}
	}
	return pre
}

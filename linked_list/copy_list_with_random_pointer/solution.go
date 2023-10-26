package copy_list_with_random_pointer

// 138. 复制带随机指针的链表 https://leetcode.cn/problems/copy-list-with-random-pointer/
// 剑指 Offer 35. 复杂链表的复制 https://leetcode.cn/problems/fu-za-lian-biao-de-fu-zhi-lcof/
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil {
		newNode := &Node{Val: cur.Val, Next: cur.Next}
		cur.Next = newNode
		cur = cur.Next.Next
	}
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	cur = head
	ret := cur.Next //用完得复原
	for cur != nil {
		temp := cur.Next
		if cur.Next != nil {
			cur.Next = cur.Next.Next
		}
		cur = temp
	}
	return ret
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

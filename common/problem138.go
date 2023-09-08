package common

// https://leetcode-cn.com/problems/copy-list-with-random-pointer/
func copyRandomList(head *Node2) *Node2 {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil {
		newNode := &Node2{Val: cur.Val, Next: cur.Next}
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

// Node2 Node2
type Node2 struct {
	Val    int
	Next   *Node2
	Random *Node2
}

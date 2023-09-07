package gosolutions

// https://leetcode-cn.com/problems/merge-k-sorted-lists/
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

/* 优先队列，堆排序
func mergeKLists(lists []*ListNode) *ListNode {
    pre:=&ListNode{}
    ret:=pre
    var node []*ListNode
    for i:=0;i<len(lists);i++ {
        if lists[i]!=nil {
            node=append(node,lists[i])
        }
    }
    for len(node)>0 {
        heapify(node)
        temp:=node[0]
        if temp.Next!=nil {node[0]=temp.Next}else{
            node=node[1:]
        }
        pre.Next=temp
        pre=pre.Next
    }
    return ret.Next
}

func heapify(node []*ListNode){
    if len(node)<2 {return }
    last:=len(node)/2-1
    for i:=last;i>=0;i-- {
        left:=2*i+1
        right:=2*i+2
        min:=i
        if left<len(node) && node[left].Val<node[min].Val {
            min=left
        }
        if right<len(node) && node[right].Val<node[min].Val {
            min=right
        }
        node[i],node[min]=node[min],node[i]
    }
}
*/

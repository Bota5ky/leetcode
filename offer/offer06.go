package offer

// https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
func reversePrint(head *ListNode) []int {
	var ret []int
	for head != nil {
		ret = append([]int{head.Val}, ret...)
		head = head.Next
	}
	return ret
}

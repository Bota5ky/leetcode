package kth_node_from_end_of_list_lcci

import . "leetcode/_model"

// 面试题 02.02. 返回倒数第 k 个节点 https://leetcode.cn/problems/kth-node-from-end-of-list-lcci/
// LCR 140. 训练计划 II https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
func kthToLast(head *ListNode, k int) int {
	rear := head
	for rear != nil && k > 0 {
		rear = rear.Next
		k--
	}
	for rear != nil {
		head = head.Next
		rear = rear.Next
	}
	return head.Val
}

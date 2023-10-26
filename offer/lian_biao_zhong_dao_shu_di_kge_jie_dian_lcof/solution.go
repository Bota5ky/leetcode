package lian_biao_zhong_dao_shu_di_kge_jie_dian_lcof

import . "leetcode/_model"

// 剑指 Offer 22. 链表中倒数第k个节点 https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
func getKthFromEnd(head *ListNode, k int) *ListNode {
	rear := head
	for ; rear != nil && k > 0; k-- {
		rear = rear.Next
	}
	for rear != nil {
		rear = rear.Next
		head = head.Next
	}
	return head
}

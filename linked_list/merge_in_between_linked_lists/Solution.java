package linked_list.merge_in_between_linked_lists;

import _model.ListNode;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/merge-in-between-linked-lists/">1669. 合并两个链表</a>
 * @since 2024-02-07 16:59:08
 */
class Solution {
    public ListNode mergeInBetween(ListNode list1, int a, int b, ListNode list2) {
        ListNode list2head = list2;
        while (list2.next != null) {
            list2 = list2.next;
        }
        ListNode list2tail = list2;
        ListNode pre = new ListNode(-1, list1);
        ListNode ret = pre;
        int i = 0;
        while (i != a) {
            i++;
            pre = pre.next;
        }
        ListNode newPre = new ListNode(-1, pre.next);
        pre.next = list2head;
        while (i != b) {
            i++;
            newPre = newPre.next;
        }
        list2tail.next = newPre.next.next;
        return ret.next;
    }
}

/*
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */

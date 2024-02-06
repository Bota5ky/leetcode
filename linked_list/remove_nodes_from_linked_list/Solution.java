package linked_list.remove_nodes_from_linked_list;

import _model.ListNode;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/remove-nodes-from-linked-list/">2487. 从链表中移除节点</a>
 * @since 2024-02-04 17:53
 */
class Solution {
    public ListNode removeNodes(ListNode head) {
        ListNode reversedNode = reverseNodes(head);
        int max = reversedNode.val;
        ListNode first = reversedNode;
        while (reversedNode.next != null) {
            ListNode nextNode = reversedNode.next;
            if (nextNode.val < max) {
                reversedNode.next = reversedNode.next.next;
            } else {
                max = nextNode.val;
                reversedNode = reversedNode.next;
            }
        }
        return reverseNodes(first);
    }

    private ListNode reverseNodes(ListNode head) {
        ListNode pre = null;
        while (head != null) {
            ListNode next = head.next;
            head.next = pre;
            pre = head;
            head = next;
        }
        return pre;
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

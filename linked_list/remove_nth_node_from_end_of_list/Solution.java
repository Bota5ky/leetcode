package linked_list.remove_nth_node_from_end_of_list;

import _model.ListNode;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/remove-nth-node-from-end-of-list/">19. 删除链表的倒数第 N 个结点</a>
 * @link <a href="https://leetcode.cn/problems/SLwz0R/">LCR 021. 删除链表的倒数第 N 个结点</a>
 * @since 2023-11-23 21:36
 */
class Solution {
    public ListNode removeNthFromEnd(ListNode head, int n) {
        ListNode newHead = new ListNode(-1, head);
        ListNode node1 = newHead;
        ListNode node2 = newHead;
        for (int i = 0; i < n; i++) {
            node2 = node2.next;
        }
        while (node2.next != null) {
            node1 = node1.next;
            node2 = node2.next;
        }
        node1.next = node1.next.next;
        return newHead.next;
    }
}

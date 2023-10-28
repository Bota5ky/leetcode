package linked_list.linked_list_cycle_ii;

import _model.ListNode;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/c32eOV/">LCR 022. 环形链表 II</a>
 * @since 2023-10-28 13:48
 */
public class Solution {
    public ListNode detectCycle(ListNode head) {
        var fast = head;
        var low = head;
        do {
            if (fast == null || fast.next == null) {
                return null;
            }
            fast = fast.next.next;
            low = low.next;
        } while (fast != low);
        fast = head;
        while (fast != low) {
            fast = fast.next;
            low = low.next;
        }
        return fast;
    }
}

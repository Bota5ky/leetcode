package linked_list.linked_list_components;

import _model.ListNode;

import java.util.Arrays;
import java.util.Set;
import java.util.stream.Collectors;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/linked-list-components/">817. 链表组件</a>
 * @since 2024-02-07 15:42
 */
class Solution {
    public int numComponents(ListNode head, int[] nums) {
        Set<Integer> set = Arrays.stream(nums).boxed().collect(Collectors.toSet());
        int count = 0;
        boolean flag = false;
        while (head != null) {
            if (set.contains(head.val)) {
                flag = true;
            } else {
                if (flag) {
                    count++;
                }
                flag = false;
            }
            head = head.next;
        }
        return count + (flag ? 1 : 0);
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

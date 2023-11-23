package _model;

public class ListNode {
    public int val;
    public ListNode next;

    public ListNode(int val) {
        this.val = val;
    }

    public ListNode(int val, ListNode next) {
        this.val = val;
        this.next = next;
    }

    public static ListNode create(int[] nums) {
        ListNode head = new ListNode(-1);
        ListNode node = head;
        for (int num : nums) {
            node.next = new ListNode(num);
            node = node.next;
        }
        return head.next;
    }
}

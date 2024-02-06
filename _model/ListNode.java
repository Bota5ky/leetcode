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

    public static ListNode build(int[] nums) {
        ListNode head = new ListNode(-1);
        ListNode node = head;
        for (int num : nums) {
            node.next = new ListNode(num);
            node = node.next;
        }
        return head.next;
    }

    public void print() {
        ListNode node = this;
        while (true) {
            if (node == null) {
                System.out.print("null");
                break;
            } else {
                System.out.printf("%d -> ", node.val);
                node = node.next;
            }
        }
    }
}

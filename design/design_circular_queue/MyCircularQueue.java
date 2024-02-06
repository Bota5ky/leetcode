package design.design_circular_queue;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/design-circular-queue/">622. 设计循环队列</a>
 * @since 2024-01-15 16:45
 */
class MyCircularQueue {
    private final int capacity;
    private int stored = 0;
    private final Node head;
    private final Node tail;

    public MyCircularQueue(int k) {
        capacity = k;
        head = new Node(null, null,-1);
        tail = new Node(head, null,-1);
        head.pre = tail;
    }

    public boolean enQueue(int value) {
        if (stored >= capacity) {
            return false;
        }
        Node nextNode = tail.next;
        Node insert = new Node(nextNode, tail, value);
        nextNode.pre = insert;
        tail.next = insert;
        stored++;
        return true;
    }

    public boolean deQueue() {
        if (stored == 0) {
            return false;
        }
        Node delete = head.pre;
        Node behind = delete.pre;
        behind.next = head;
        head.pre = behind;
        stored--;
        return true;
    }

    public int Front() {
        return head.pre.val;
    }

    public int Rear() {
        return tail.next.val;
    }

    public boolean isEmpty() {
        return stored == 0;
    }

    public boolean isFull() {
        return stored == capacity;
    }

    static class Node {
        Node next;
        Node pre;
        int val;

        public Node(Node next,Node pre, int val) {
            this.next = next;
            this.pre = pre;
            this.val = val;
        }
    }
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * MyCircularQueue obj = new MyCircularQueue(k);
 * boolean param_1 = obj.enQueue(value);
 * boolean param_2 = obj.deQueue();
 * int param_3 = obj.Front();
 * int param_4 = obj.Rear();
 * boolean param_5 = obj.isEmpty();
 * boolean param_6 = obj.isFull();
 */

package design.lfu_cache;

import java.util.HashMap;
import java.util.Map;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/lfu-cache/">460. LFU 缓存</a>
 * @since 2023-09-25 15:29
 */
class LFUCache {
    private final DoubleLinkedList tailList;
    private final Map<Integer, Node> cache;
    private final int capacity;

    public LFUCache(int capacity) {
        this.capacity = capacity;
        var headList = new DoubleLinkedList(Integer.MAX_VALUE);
        this.tailList = new DoubleLinkedList(0);
        headList.lessFreq = this.tailList;
        this.tailList.moreFreq = headList;
        this.cache = new HashMap<>(capacity);
    }

    public int get(int key) {
        if (cache.containsKey(key)) {
            var node = cache.get(key);
            var list = node.doubleLinkedList;
            var nodeFreq = list.frequency + 1;
            list.removeNode(node);
            if (list.moreFreq.frequency > nodeFreq) {
                var newAddedList = new DoubleLinkedList(nodeFreq);
                newAddedList.addNode(node);
                addDoubleLinkedListAt(list, newAddedList);
            } else {
                list.moreFreq.addNode(node);
            }
            if (list.head.next == list.tail) {
                removeDoubleLinkedList(list);
            }
            return node.val;
        }
        return -1;
    }

    public void put(int key, int value) {
        if (cache.containsKey(key)) {
            var node = cache.get(key);
            node.val = value;
            get(key);
        } else {
            var node = new Node(key, value);
            var leastFreqList = tailList.moreFreq;
            if (capacity == cache.size()) {
                var removed = leastFreqList.tail.pre;
                cache.remove(removed.key);
                leastFreqList.removeNode(removed);
                if (leastFreqList.head.next == leastFreqList.tail) {
                    removeDoubleLinkedList(leastFreqList);
                }
            }
            leastFreqList = tailList.moreFreq;
            if (leastFreqList.frequency != 1) {
                var newAddedList = new DoubleLinkedList(1);
                newAddedList.addNode(node);
                addDoubleLinkedListAt(tailList, newAddedList);
            } else {
                leastFreqList.addNode(node);
            }
            cache.put(key, node);
        }
    }

    private void addDoubleLinkedListAt(DoubleLinkedList target, DoubleLinkedList newAddedList) {
        var moreFreqList = target.moreFreq;
        moreFreqList.lessFreq = newAddedList;
        newAddedList.moreFreq = moreFreqList;
        newAddedList.lessFreq = target;
        target.moreFreq = newAddedList;
    }

    private void removeDoubleLinkedList(DoubleLinkedList list) {
        var moreFreqList = list.moreFreq;
        var lessFreqList = list.lessFreq;
        moreFreqList.lessFreq = lessFreqList;
        lessFreqList.moreFreq = moreFreqList;
    }

    private static class DoubleLinkedList {
        int frequency;
        Node head;
        Node tail;
        DoubleLinkedList moreFreq;
        DoubleLinkedList lessFreq;

        DoubleLinkedList(int frequency) {
            this.frequency = frequency;
            this.head = new Node();
            this.tail = new Node();
            this.head.next = this.tail;
            this.tail.pre = this.head;
        }

        void addNode(Node node) {
            node.doubleLinkedList = this;
            var preHead = head.next;
            preHead.pre = node;
            node.pre = head;
            head.next = node;
            node.next = preHead;
        }

        void removeNode(Node node) {
            var nextNode = node.next;
            var preNode = node.pre;
            nextNode.pre = preNode;
            preNode.next = nextNode;
        }
    }

    private static class Node {
        int key;
        int val;
        Node pre;
        Node next;
        DoubleLinkedList doubleLinkedList;

        Node() {
        }

        Node(int key, int val) {
            this.key = key;
            this.val = val;
        }
    }
}

/*
 * Your LFUCache object will be instantiated and called as such:
 * LFUCache obj = new LFUCache(capacity);
 * int param_1 = obj.get(key);
 * obj.put(key,value);
 */

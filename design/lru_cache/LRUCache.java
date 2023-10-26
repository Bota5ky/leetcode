package design.lru_cache;

import java.util.HashMap;
import java.util.Map;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/lru-cache/">146. LRU 缓存</a>
 * @link <a href="https://leetcode.cn/problems/lru-cache-lcci/">面试题 16.25. LRU 缓存</a>
 * @since 2023-09-24 11:29
 */
class LRUCache {
    private final int capacity;
    private final Map<Integer, Node> cache;
    private final Node head;
    private final Node tail;

    private static class Node {
        Node pre;
        Node next;
        int key;
        int val;
        Node() {
        }
        Node(int key, int val) {
            this.key = key;
            this.val = val;
        }
    }

    public LRUCache(int capacity) {
        this.capacity = capacity;
        this.cache = new HashMap<>();
        this.head = new Node();
        this.tail = new Node();
        this.tail.next = this.head;
        this.head.pre = this.tail;
    }

    public int get(int key) {
        if (cache.containsKey(key)) {
            var getNode = cache.get(key);
            deleteNode(getNode);
            insertNode(getNode);
            return getNode.val;
        }
        return -1;
    }

    public void put(int key, int value) {
        if (cache.containsKey(key)) {
            var getNode = cache.get(key);
            deleteNode(getNode);
            insertNode(getNode);
            getNode.val = value;
        } else {
            if (cache.size() == this.capacity) {
                deleteNode(this.tail.next);
            }
            insertNode(new Node(key, value));
        }
    }

    private void insertNode(Node node) {
        cache.put(node.key, node);
        var preHead = this.head.pre;
        preHead.next = node;
        node.pre = preHead;
        node.next = this.head;
        this.head.pre = node;
    }

    private void deleteNode(Node node) {
        cache.remove(node.key);
        var preNode = node.pre;
        preNode.next = node.next;
        node.next.pre = preNode;
    }
}

/*
JDK自带实现
class LRUCache extends LinkedHashMap<Integer, Integer> {
    private final int capacity;

    public LRUCache(int capacity) {
        super(capacity, 0.75F, true);
        this.capacity = capacity;
    }

    public int get(int key) {
        return super.getOrDefault(key, -1);
    }

    public void put(int key, int value) {
        super.put(key, value);
    }

    @Override
    protected boolean removeEldestEntry(Map.Entry<Integer, Integer> eldest) {
        return size() > capacity;
    }
}
*/

/*
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache obj = new LRUCache(capacity);
 * int param_1 = obj.get(key);
 * obj.put(key,value);
 */

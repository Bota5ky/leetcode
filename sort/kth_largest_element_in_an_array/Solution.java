package sort.kth_largest_element_in_an_array;

import java.util.PriorityQueue;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/kth-largest-element-in-an-array/">215. 数组中的第K个最大元素</a>
 * @since 2023-11-03 10:21
 */
class Solution {
    public int findKthLargest(int[] nums, int k) {
        var heap = new PriorityQueue<Integer>();
        for (int num : nums) {
            if (heap.size() < k) {
                heap.offer(num);
            } else if (heap.peek() < num) {
                heap.poll();
                heap.offer(num);
            }
        }
        return heap.peek();
    }
}

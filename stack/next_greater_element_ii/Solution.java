package stack.next_greater_element_ii;

import java.util.Arrays;
import java.util.Comparator;
import java.util.PriorityQueue;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/next-greater-element-ii/">503. 下一个更大元素 II</a>
 * @since 2024-06-24 08:22
 */
class Solution {
    public int[] nextGreaterElements(int[] nums) {
        int[] bigger = new int[nums.length];
        Arrays.fill(bigger, -1);
        // [value,position]
        PriorityQueue<int[]> queue = new PriorityQueue<>(Comparator.comparingInt(l -> l[0]));
        for (int i = 0; i < nums.length; i++) {
            while (!queue.isEmpty() && nums[i] > queue.peek()[0]) {
                int[] poll = queue.poll();
                bigger[poll[1]] = nums[i];
            }
            queue.add(new int[]{nums[i], i});
        }
        for (int i = 0; i < nums.length && !queue.isEmpty(); i++) {
            while (!queue.isEmpty() && nums[i] > queue.peek()[0]) {
                int[] poll = queue.poll();
                bigger[poll[1]] = nums[i];
            }
            queue.add(new int[]{nums[i], i});
        }
        return bigger;
    }
}

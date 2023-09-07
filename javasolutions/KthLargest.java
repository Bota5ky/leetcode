package javasolutions;

import java.util.ArrayList;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/kth-largest-element-in-a-stream/">703. 数据流中的第 K 大元素</a>
 * @since 2023-09-07 14:12
 */
class KthLargest {
    private int limit;
    private ArrayList<Integer> stack;

    public KthLargest(int k, int[] nums) {
        limit = k;
        stack = new ArrayList<>();
        for (int i = 0; i < nums.length && i < k; i++) {
            stack.add(nums[i]);
        }
        for (int i = k; i < nums.length; i++) {
            heapify(stack);
            if (stack.get(0) < nums[i]) {
                stack.set(0, nums[i]);
            }
        }
        heapify(stack);
    }

    public int add(int val) {
        if (stack.size() < limit) {
            stack.add(val);
        } else {
            if (stack.get(0) < val) {
                stack.set(0, val);
            } else {
                return stack.get(0);
            }
        }
        heapify(stack);
        return stack.get(0);
    }

    private void heapify(ArrayList<Integer> stack) {
        int last = stack.size() / 2 - 1;
        for (int i = last; i >= 0; i--) {
            int min = i, left = 2 * i + 1, right = 2 * i + 2;
            if (stack.get(left) < stack.get(min)) {
                min = left;
            }
            if (right < stack.size() && stack.get(right) < stack.get(min)) {
                min = right;
            }
            int temp = stack.get(min);
            stack.set(min, stack.get(i));
            stack.set(i, temp);
        }
    }
}

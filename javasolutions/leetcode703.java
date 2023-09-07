package javasolutions;

import java.util.ArrayList;

//https://leetcode-cn.com/problems/kth-largest-element-in-a-stream/
// public class leetcode703 {
//     private PriorityQueue<Integer> queue;
//     private int limit;

//     public leetcode703(int k, int[] nums) {
//         limit = k;
//         queue = new PriorityQueue<>(k);
//         for (int num : nums) {
//             add(num);
//         }
//     }

//     public int add(int val) {
//         if (queue.size() < limit) {
//             queue.add(val);
//         } else if (val > queue.peek()) {
//             queue.poll();
//             queue.add(val);
//         }

//         return queue.peek();
//     }
// }

public class leetcode703 {
    private int limit;
    private ArrayList<Integer> stack;

    public leetcode703(int k, int[] nums) {
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
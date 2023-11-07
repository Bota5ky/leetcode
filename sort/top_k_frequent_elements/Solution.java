package sort.top_k_frequent_elements;

import java.util.Comparator;
import java.util.HashMap;
import java.util.Map;
import java.util.PriorityQueue;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/g5c51o/">LCR 060. 前 K 个高频元素</a>
 * @since 2023-11-07 15:03
 */
class Solution {
    public int[] topKFrequent(int[] nums, int k) {
        Map<Integer, Integer> countMap = new HashMap<>();
        for (int num : nums) {
            int count = countMap.getOrDefault(num, 0);
            countMap.put(num, ++count);
        }
        PriorityQueue<int[]> p = new PriorityQueue<>(Comparator.comparingInt(l -> l[1]));
        for (Map.Entry<Integer, Integer> entry : countMap.entrySet()) {
            int num = entry.getKey();
            int count = entry.getValue();
            if (p.size() < k) {
                p.offer(new int[]{num, count});
                continue;
            }
            if (p.peek()[1] < count){
                p.poll();
                p.offer(new int[]{num, count});
            }
        }
        int[] res = new int[k];
        for (int i = 0; i < k; ++i) {
            res[i] = p.poll()[0];
        }
        return res;
    }
}

package array.minimum_seconds_to_equalize_a_circular_array;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/minimum-seconds-to-equalize-a-circular-array/">2808. 使循环数组所有元素相等的最少秒数</a>
 * @since 2024-01-31 16:46
 */
class Solution {
    public int minimumSeconds(List<Integer> nums) {
        Map<Integer, Integer> fp = new HashMap<>();
        Map<Integer, Integer> lp = new HashMap<>();
        Map<Integer, Integer> gap = new HashMap<>();
        for (int i = 0; i < nums.size(); i++) {
            Integer val = nums.get(i);
            fp.putIfAbsent(val, i);
            if (!lp.containsKey(val)) {
                gap.put(val, 0);
            } else {
                int distance = i - lp.get(val) - 1;
                if (gap.get(val) < distance) {
                    gap.put(val, distance);
                }
            }
            lp.put(val, i);
        }
        for (Integer key : fp.keySet()) {
            int inverseDistance = nums.size() - (lp.get(key) - fp.get(key)) - 1;
            if (!gap.containsKey(key) || gap.get(key) < inverseDistance) {
                gap.put(key, inverseDistance);
            }
        }
        int result = nums.size() - 1;
        for (Integer value : gap.values()) {
            if (result > value) {
                result = value;
            }
        }
        return (result + 1) / 2;
    }
}

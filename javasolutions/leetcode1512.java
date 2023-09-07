package javasolutions;

import java.util.HashMap;
import java.util.Map;

//https://leetcode-cn.com/problems/number-of-good-pairs/
public class leetcode1512 {
    public int numIdenticalPairs(int[] nums) {
        // 1+..+n-1 n(n-1)/2
        int sum = 0;
        Map<Integer, Integer> map = new HashMap<>();
        for (int num : nums) {
            if (map.containsKey(num)) {
                sum += map.get(num);
                map.put(num, map.get(num) + 1);
            } else {
                map.put(num, 1);
            }
        }
        return sum;
    }
}

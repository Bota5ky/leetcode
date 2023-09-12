package number_of_good_pairs

import java.util.HashMap;
import java.util.Map;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/number-of-good-pairs/">1512. 好数对的数目</a>
 * @since 2023-09-07 14:27
 */
class Problem1512 {
    public int numIdenticalPairs(int[] nums) {
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

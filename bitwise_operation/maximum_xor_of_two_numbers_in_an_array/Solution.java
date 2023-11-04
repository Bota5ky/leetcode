package bitwise_operation.maximum_xor_of_two_numbers_in_an_array;

import java.util.HashSet;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/maximum-xor-of-two-numbers-in-an-array/">421. 数组中两个数的最大异或值</a>
 * @since 2023-11-04 18:56
 */
class Solution {
    public int findMaximumXOR(int[] nums) {
        int res = 0;
        int mask = 0;
        for (int i = 30; i >= 0; i--) {
            mask = mask ^ (1 << i);
            HashSet<Integer> set = new HashSet<>();
            for (int num : nums) {
                set.add(num & mask);
            }
            int temp = res ^ (1 << i);
            for (Integer prefix : set) {
                if (set.contains(prefix ^ temp)) {
                    res = temp;
                    break;
                }
            }
        }
        return res;
    }
}

package bitwise_operation.single_number_ii;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/single-number-ii/">137. 只出现一次的数字 II</a>
 * @since 2023-10-15 09:21
 */
class Solution {
    public int singleNumber(int[] nums) {
        int ones = 0;
        int twos = 0;
        for (int num : nums) {
            ones = ~twos & (ones ^ num);
            twos = ~ones & (twos ^ num);
        }
        return ones;
    }
}
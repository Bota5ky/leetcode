package common.question_151_to_200.majority_element;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/majority-element/">169. 多数元素</a>
 * @since 2023-09-15 16:30
 */
class Solution {
    public int majorityElement(int[] nums) {
        var hold = Integer.MAX_VALUE;
        var count = 0;
        for (int num : nums) {
            if (count == 0 || hold == num) {
                count++;
                hold = num;
            } else {
                count--;
            }
        }
        return hold;
    }
}

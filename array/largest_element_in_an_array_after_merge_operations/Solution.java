package array.largest_element_in_an_array_after_merge_operations;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/largest-element-in-an-array-after-merge-operations/">2789. 合并后数组中的最大元素</a>
 * @since 2024-03-14 22:48
 */
class Solution {
    public long maxArrayValue(int[] nums) {
        long max = nums[nums.length - 1];
        for (int i = nums.length - 1; i > 0; i--) {
            max = nums[i - 1] <= max ? nums[i - 1] + max : nums[i - 1];
        }
        return max;
    }
}

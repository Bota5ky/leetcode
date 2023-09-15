package common.question_51_to_100.remove_duplicates_from_sorted_array_ii;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/">80. 删除有序数组中的重复项 II</a>
 * @since 2023-09-15 16:00
 */
class Solution {
    public int removeDuplicates(int[] nums) {
        if (nums.length < 3) {
            return nums.length;
        }
        var p = 1;
        for (int i = 2; i < nums.length; i++) {
            if (nums[i] != nums[p - 1]) {
                nums[++p] = nums[i];
            }
        }
        return p + 1;
    }
}

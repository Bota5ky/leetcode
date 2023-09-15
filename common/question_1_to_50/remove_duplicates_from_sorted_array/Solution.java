package common.question_1_to_50.remove_duplicates_from_sorted_array;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/remove-duplicates-from-sorted-array/">26. 删除有序数组中的重复项</a>
 * @since 2023-09-15 15:53
 */
class Solution {
    public int removeDuplicates(int[] nums) {
        var p = 0;
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] != nums[p]) {
                nums[++p] = nums[i];
            }
        }
        return p + 1;
    }
}

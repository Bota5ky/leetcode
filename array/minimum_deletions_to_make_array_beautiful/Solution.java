package array.minimum_deletions_to_make_array_beautiful;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/minimum-deletions-to-make-array-beautiful/">2216. 美化数组的最少删除数</a>
 * @since 2023-11-21 19:04
 */
class Solution {
    public int minDeletion(int[] nums) {
        int removed = 0;
        int pos = 0;
        for (int i = 0; i < nums.length - 1; i++) {
            if (pos % 2 == 0 && nums[i] == nums[i + 1]) {
                removed++;
                pos--;
            }
            pos++;
        }
        if (pos % 2 == 0) {
            removed++;
        }
        return removed;
    }
}

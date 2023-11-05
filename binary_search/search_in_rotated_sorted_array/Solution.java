package binary_search.search_in_rotated_sorted_array;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/search-in-rotated-sorted-array/">33. 搜索旋转排序数组</a>
 * @since 2023-11-05 18:12
 */
class Solution {
    public int search(int[] nums, int target) {
        for (int i = 0, j = nums.length - 1; i <= j; ) {
            int mid = (i + j) / 2;
            if (nums[mid] == target) {
                return mid;
            }
            if (nums[mid] >= nums[i]) {
                if (nums[i] <= target && target < nums[mid]) {
                    j = mid - 1;
                } else {
                    i = mid + 1;
                }
            } else {
                if (nums[j] >= target && target > nums[mid]) {
                    i = mid + 1;
                } else {
                    j = mid - 1;
                }
            }
        }
        return -1;
    }
}

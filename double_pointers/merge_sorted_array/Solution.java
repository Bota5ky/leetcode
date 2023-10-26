package double_pointers.merge_sorted_array;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/merge-sorted-array/">88. 合并两个有序数组</a>
 * @since 2023-09-15 15:20
 */
class Solution {
    public void merge(int[] nums1, int m, int[] nums2, int n) {
        for (int i = m + n - 1; i >= 0 && n > 0; i--) {
            if (m > 0 && nums1[m - 1] >= nums2[n - 1]) {
                nums1[i] = nums1[--m];
            } else {
                nums1[i] = nums2[--n];
            }
        }
    }
}

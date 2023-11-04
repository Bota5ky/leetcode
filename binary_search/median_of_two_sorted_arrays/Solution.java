package binary_search.median_of_two_sorted_arrays;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/median-of-two-sorted-arrays/">4. 寻找两个正序数组的中位数</a>
 * @since 2023-11-03 14:47
 */
class Solution {
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int len1 = nums1.length;
        int len2 = nums2.length;
        int odd = findKth(nums1, 0, len1 - 1, nums2, 0, len2 - 1, (len1 + len2 + 1) / 2);
        int even = findKth(nums1, 0, len1 - 1, nums2, 0, len2 - 1, (len1 + len2 + 2) / 2);
        return (odd + even) / 2.0;
    }

    private int findKth(int[] nums1, int start1, int end1, int[] nums2, int start2, int end2, int k) {
        int len1 = end1 - start1 + 1;
        int len2 = end2 - start2 + 1;
        if (len1 > len2) {
            return findKth(nums2, start2, end2, nums1, start1, end1, k);
        }
        if (len1 == 0) {
            return nums2[start2 + k - 1];
        }
        if (k == 1) {
            return Math.min(nums1[start1], nums2[start2]);
        }
        int i = start1 + Math.min(len1, k / 2) - 1;
        int j = start2 + Math.min(len2, k / 2) - 1;
        if (nums1[i] > nums2[j]) {
            return findKth(nums1, start1, end1, nums2, j + 1, end2, k - (j + 1 - start2));
        } else {
            return findKth(nums1, i + 1, end1, nums2, start2, end2, k - (i + 1 - start1));
        }
    }
}

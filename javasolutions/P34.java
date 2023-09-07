package javasolutions;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/">34. 在排序数组中查找元素的第一个和最后一个位置</a>
 * @since 2023-09-07 10:08
 */
class P34 {
    public int[] searchRange(int[] nums, int target) {
        int[] res = {-1, -1};
        int i = 0, j = nums.length - 1;
        while (i < j) {
            int mid = (i + j) / 2;
            if (nums[mid] >= target) {
                j = mid;
            } else {
                i = mid + 1;
            }
        }
        int left = i;
        if (j < 0 || nums[left] != target) {
            return res;
        }
        i = 0;
        j = nums.length - 1;
        while (i < j) {
            int mid = (i + j) / 2;
            if (nums[mid] > target) {
                j = mid;
            } else {
                i = mid + 1;
            }
        }
        int right = i;
        if (nums[right] == target) {
            right++;
        }
        res[0] = left;
        res[1] = right - 1;
        return res;
    }
}

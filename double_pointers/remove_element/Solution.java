package double_pointers.remove_element;

import java.util.Arrays;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/remove-element/">27. 移除元素</a>
 * @since 2023-09-15 15:42
 */
class Solution {
    public int removeElement(int[] nums, int val) {
        var last = nums.length - 1;
        for (int i = 0; i < nums.length && i <= last; i++) {
            if (nums[i] == val) {
                nums[i--] = nums[last--];
            }
        }
        nums = Arrays.copyOf(nums, last + 1);
        return nums.length;
    }
}
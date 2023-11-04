package array.product_of_array_except_self;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/product-of-array-except-self/">238. 除自身以外数组的乘积</a>
 * @since 2023-09-18 15:00
 */
class Solution {
    public int[] productExceptSelf(int[] nums) {
        int[] res = new int[nums.length];
        res[0] = 1;
        for (int i = 1, pre = nums[0]; i < nums.length; i++) {
            res[i] = pre;
            pre *= nums[i];
        }
        for (int i = nums.length - 2, pre = nums[nums.length - 1]; i >= 0 ; i--) {
            res[i] *= pre;
            pre *= nums[i];
        }
        return res;
    }
}

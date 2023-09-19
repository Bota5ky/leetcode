package common.question_1350_to_inf.house_robber_iv;

import java.util.Arrays;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/house-robber-iv/">2560. 打家劫舍 IV</a>
 * @since 2023-09-19 08:53
 */
class Solution {
    public int minCapability(int[] nums, int k) {
        var r = Arrays.stream(nums).max().getAsInt();
        var l = Arrays.stream(nums).min().getAsInt();
        while (l <= r) {
            int cnt = 0;
            int mid = (l + r) / 2;
            boolean visited = false;
            for (int num : nums) {
                if (num <= mid && !visited) {
                    cnt++;
                    visited = true;
                } else {
                    visited = false;
                }
            }
            if (cnt >= k) {
                r = mid - 1;
            } else {
                l = mid + 1;
            }
        }
        return l;
    }
}

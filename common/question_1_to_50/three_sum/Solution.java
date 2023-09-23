package common.question_1_to_50.three_sum;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/3sum/">15. 三数之和</a>
 * @since 2023-09-23 19:56
 */
class Solution {
    public List<List<Integer>> threeSum(int[] nums) {
        Arrays.sort(nums);
        var result = new ArrayList<List<Integer>>();
        for (int i = 0; i < nums.length - 2; i++) {
            if (i > 0 && nums[i] == nums[i - 1]) {
                continue;
            }
            for (int j = i + 1; j < nums.length - 1; j++) {
                if (j > i + 1 && nums[j] == nums[j - 1]) {
                    continue;
                }
                var target = -nums[i] - nums[j];
                var low = j + 1;
                var high = nums.length - 1;
                if (nums[low] > target || nums[high] < target) {
                    continue;
                }
                while (low <= high) {
                    var mid = (low + high) / 2;
                    if (nums[mid] == target) {
                        var list = new ArrayList<Integer>();
                        list.add(nums[i]);
                        list.add(nums[j]);
                        list.add(nums[mid]);
                        result.add(list);
                        break;
                    } else if (nums[mid] < target) {
                        low = mid + 1;
                    } else {
                        high = mid - 1;
                    }
                }
            }
        }
        return result;
    }
}

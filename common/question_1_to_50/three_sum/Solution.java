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
        var res = new ArrayList<List<Integer>>();
        for (int i = 0; i < nums.length - 2 && nums[i] <= 0; i++) {
            if (i > 0 && nums[i] == nums[i - 1]) {
                continue;
            }
            for (int j = i + 1; j < nums.length - 1 && nums[i] + nums[j] <= 0; j++) {
                if (j > i + 1 && nums[j] == nums[j - 1]) {
                    continue;
                }
                // binarySearch
                var head = j + 1;
                var tail = nums.length - 1;
                while (head <= tail) {
                    var mid = (head + tail) / 2;
                    int sum = nums[i] + nums[j] + nums[mid];
                    if (sum == 0) {
                        res.add(List.of(nums[i], nums[j], nums[mid]));
                        break;
                    } else if (sum > 0) {
                        tail = mid - 1;
                    } else {
                        head = mid + 1;
                    }
                }
            }
        }
        return res;
    }
}

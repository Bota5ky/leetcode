package common.question_151_to_200.two_sum_ii_input_array_is_sorted;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/">167. 两数之和 II - 输入有序数组</a>
 * @since 2023-09-22 19:26
 */
class Solution {
    // binary search O(nlogn)
    public int[] twoSum(int[] numbers, int target) {
        for (int i = 0; i < numbers.length - 1; i++) {
            var low = i + 1;
            var high = numbers.length - 1;
            while (low <= high) {
                var mid = (low + high) / 2;
                if (numbers[mid] == target - numbers[i]) {
                    return new int[]{i + 1, mid + 1};
                } else if (numbers[mid] < target - numbers[i]) {
                    low = mid + 1;
                } else if (numbers[mid] > target - numbers[i]) {
                    high = mid - 1;
                }
            }
        }
        return new int[]{-1, -1};
    }
}

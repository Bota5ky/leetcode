package math.sum_multiples;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/sum-multiples/">2652. 倍数求和</a>
 * @since 2023-10-17 14:18
 */
class Solution {
    public int sumOfMultiples(int n) {
        var sum = 0;
        var prev = 0;
        int[] nums = {3, 5, 7};
        int[] gain = {1, 1, 1};
        for (; ; ) {
            var first = nums[0] * gain[0];
            var second = nums[1] * gain[1];
            var third = nums[2] * gain[2];
            if (first > n && second > n && third > n) {
                return sum + prev;
            }
            if (first <= second && first <= third) {
                if (first > prev) {
                    sum += prev;
                    prev = first;
                }
                gain[0] += 1;
            }
            if (second <= first && second <= third) {
                if (second > prev) {
                    sum += prev;
                    prev = second;
                }
                gain[1] += 1;
            }
            if (third <= first && third <= second) {
                if (third > prev) {
                    sum += prev;
                    prev = third;
                }
                gain[2] += 1;
            }
        }
    }
}

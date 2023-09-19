package common.question_101_to_150.candy;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/candy/">135. 分发糖果</a>
 * @since 2023-09-19 14:16
 */
class Solution {
    public int candy(int[] ratings) {
        var left = new int[ratings.length];
        for (int i = 0, pre = 0; i < ratings.length; i++) {
            if (i == 0 || ratings[i] > ratings[i - 1]) {
                pre++;
            } else {
                pre = 1;
            }
            left[i] = pre;
        }
        var sum = 0;
        for (int i = ratings.length - 1, pre = 0; i >= 0; i--) {
            if (i == ratings.length - 1 || ratings[i] > ratings[i + 1]) {
                pre++;
            } else {
                pre = 1;
            }
            sum += Math.max(pre, left[i]);
        }
        return sum;
    }
}

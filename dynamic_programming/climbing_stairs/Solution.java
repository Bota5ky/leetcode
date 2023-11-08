package dynamic_programming.climbing_stairs;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/climbing-stairs/">70. 爬楼梯</a>
 * @since 2023-11-08 15:16
 */
class Solution {
    public int climbStairs(int n) {
        int[] steps = new int[n + 1];
        steps[0] = 1;
        steps[1] = 1;
        for (int i = 2; i <= n; i++) {
            steps[i] = steps[i - 1] + steps[i - 2];
        }
        return steps[n];
    }
}

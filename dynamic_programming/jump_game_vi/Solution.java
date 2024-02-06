package dynamic_programming.jump_game_vi;

import java.util.Arrays;
import java.util.Deque;
import java.util.LinkedList;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/jump-game-vi/">1696. 跳跃游戏 VI</a>
 * @since 2024-02-06 11:32
 */
class Solution {
    public int maxResult(int[] nums, int k) {
        int len = nums.length;
        Deque<int[]> q = new LinkedList<>();
        if (len == 1) {
            return nums[0];
        }
        int[] dp = new int[len];
        for (int i = len - 1; i >= 0; i--) {
            dp[i] = getMax(q, i, k) + nums[i];
            add(q, i, k, dp[i]);
        }
        // dp[i] = dp[i + 1] + nums[i] ... dp[i + k] + nums[i]
        System.out.println(Arrays.toString(dp));
        return dp[0];
    }

    private int getMax(Deque<int[]> q, int i, int k) {
        while (!q.isEmpty() && q.getFirst()[1] - i > k) {
            q.removeFirst();
        }
        if (q.isEmpty()) {
            return 0;
        }
        return q.getFirst()[0];
    }

    private void add(Deque<int[]> q, int i, int k, int num) {
        while (!q.isEmpty() && (q.getLast()[0] <= num || q.getLast()[1] - i > k)) {
            q.removeLast();
        }
        q.addLast(new int[]{num, i});
    }
}

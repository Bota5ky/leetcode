package array.magic_tower_game;

import java.util.PriorityQueue;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/p0NxJO/">LCP 30. 魔塔游戏</a>
 * @since 2024-02-06 16:47
 */
class Solution {
    public int magicTower(int[] nums) {
        PriorityQueue<Integer> q = new PriorityQueue<>();
        long stored = 0;
        long sum = 1;
        int count = 0;
        for (int num : nums) {
            sum += num;
            if (num < 0) {
                q.offer(num);
            }
            while (!q.isEmpty() && sum <= 0) {
                int store = q.poll();
                stored += store;
                sum -= store;
                count++;
            }
        }
        return sum + stored > 0 ? count : -1;
    }
}

package stack.asteroid_collision;

import java.util.Stack;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/asteroid-collision/">735. 小行星碰撞</a>
 * @since 2023-11-16 19:16
 */
class Solution {
    public int[] asteroidCollision(int[] asteroids) {
        Stack<Integer> stack = new Stack<>();
        for (int i = 0; i < asteroids.length; i++) {
            int asteroid = asteroids[i];
            if (!stack.isEmpty() && stack.peek() > 0 && asteroid < 0) {
                if (-asteroid == stack.peek()) {
                    stack.pop();
                    continue;
                }
                if (-asteroid > stack.peek()) {
                    stack.pop();
                    i--;
                }
                continue;
            }
            stack.push(asteroid);
        }
        int[] res = new int[stack.size()];
        for (int i = res.length - 1; i >= 0; i--) {
            res[i] = stack.pop();
        }
        return res;
    }
}
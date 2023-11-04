package greedy.jump_game;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/jump-game/">55. 跳跃游戏</a>
 * @since 2023-09-17 17:47
 */
class Solution {
    public boolean canJump(int[] nums) {
        int i = 0;
        int j = 0;
        while(j < nums.length){
            int pre = j;
            for (int k = i; k <= j && k < nums.length; k++) {
                j = Math.max(j, k + nums[k]);
            }
            if (pre == j) {
                break;
            }
            i = pre + 1;
        }
        return j >= nums.length - 1;
    }
}

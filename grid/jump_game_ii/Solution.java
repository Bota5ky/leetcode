package grid.jump_game_ii;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/jump-game-ii/">45. 跳跃游戏 II</a>
 * @since 2023-09-18 12:47
 */
class Solution {
    public int jump(int[] nums) {
        var jumpCnt = 0;
        for (int i = 0, j = 0; ; ) {
            if (j >= nums.length - 1) {
                break;
            }
            var pre = j;
            for (int k = i; k <= pre; k++) {
                j = Math.max(j, nums[k] + k);
            }
            i = pre + 1;
            jumpCnt++;
        }
        return jumpCnt;
    }
}

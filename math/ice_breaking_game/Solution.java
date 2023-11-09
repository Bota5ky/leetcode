package math.ice_breaking_game;

import java.util.ArrayList;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/">LCR 187. 破冰游戏</a>
 * @since 2023-11-09 15:01
 */
class Solution {
    public int iceBreakingGame(int num, int target) {
        ArrayList<Integer> nums = new ArrayList<>(num);
        for (int i = 0; i < num; i++) {
            nums.add(i);
        }
        int i = 0;
        while(nums.size() > 1){
            i = (i + target - 1) % nums.size();
            nums.remove(i);
        }
        return nums.get(0);
    }
}

package array.find_champion_i;

import java.util.HashSet;
import java.util.Set;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/find-champion-i/">2923. 找到冠军 I</a>
 * @since 2024-04-12 22:58
 */
class Solution {
    public int findChampion(int[][] grid) {
        Set<Integer> set = new HashSet<>();
        for (int i = 0; i < grid.length; i++) {
            set.add(i);
        }
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[i].length; j++) {
                if (i == j) {
                    continue;
                }
                if (grid[i][j] == 1) {
                    set.remove(j);
                } else {
                    set.remove(i);
                }
            }
        }
        return set.iterator().next();
    }
}

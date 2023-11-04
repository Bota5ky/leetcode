package sort.relative_ranks;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/relative-ranks/">506. 相对名次</a>
 * @since 2023-09-07 14:07
 */
class Solution {
    public String[] findRelativeRanks(int[] score) {
        Map<Integer, Integer> originPosition = new HashMap<>();
        for (int i = 0; i < score.length; i++) {
            originPosition.put(score[i], i);
        }
        String[] rank = new String[score.length];
        Arrays.sort(score);
        for (int i = 0; i < score.length; i++) {
            int rankPosition = score.length - i;
            if (rankPosition <= 3) {
                switch (rankPosition) {
                    case 1:
                        rank[originPosition.get(score[i])] = "Gold Medal";
                        break;
                    case 2:
                        rank[originPosition.get(score[i])] = "Silver Medal";
                        break;
                    case 3:
                        rank[originPosition.get(score[i])] = "Bronze Medal";
                        break;
                }
            } else {
                rank[originPosition.get(score[i])] = String.valueOf(rankPosition);
            }
        }
        return rank;
    }
}

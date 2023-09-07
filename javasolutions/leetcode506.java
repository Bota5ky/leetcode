package javasolutions;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class leetcode506 {
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
                    case 1 -> rank[originPosition.get(score[i])] = "Gold Medal";
                    case 2 -> rank[originPosition.get(score[i])] = "Silver Medal";
                    case 3 -> rank[originPosition.get(score[i])] = "Bronze Medal";
                }
            } else {
                rank[originPosition.get(score[i])] = String.valueOf(rankPosition);
            }
        }
        return rank;
    }
}

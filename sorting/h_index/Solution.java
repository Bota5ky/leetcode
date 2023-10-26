package sorting.h_index;

import java.util.Arrays;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/h-index/">274. H 指数</a>
 * @since 2023-09-18 13:09
 */
class Solution {
    public int hIndex(int[] citations) {
        Arrays.sort(citations);
        var h = 0;
        for (int i = citations.length - 1; i >= 0 && citations[i] > h; i--) {
            h++;
        }
        return h;
    }
}

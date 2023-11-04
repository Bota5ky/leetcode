package sort.h_index_ii;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/h-index-ii/">275. H 指数 II</a>
 * @since 2023-10-30 10:27
 */
class Solution {
    public int hIndex(int[] citations) {
        int h = 0;
        for (int i = 0, j = citations.length - 1; i <= j;) {
            int mid = (i + j) / 2;
            if (citations[mid] >= citations.length - mid) {
                h = Math.max(h, citations.length - mid);
                j = mid - 1;
            } else {
                i = mid + 1;
            }
        }
        return h;
    }
}

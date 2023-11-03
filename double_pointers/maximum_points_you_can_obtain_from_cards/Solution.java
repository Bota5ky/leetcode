package double_pointers.maximum_points_you_can_obtain_from_cards;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/maximum-points-you-can-obtain-from-cards/">1423. 可获得的最大点数</a>
 * @since 2023-11-03 20:57
 */
class Solution {
    public int maxScore(int[] cardPoints, int k) {
        int max = 0;
        for (int i = 0; i < k; i++) {
            max += cardPoints[i];
        }
        var sum = max;
        for (int i = k - 1, j = cardPoints.length - 1; i >= 0; i--) {
            sum = sum - cardPoints[i] + cardPoints[j--];
            max = Math.max(max, sum);
        }
        return max;
    }
}

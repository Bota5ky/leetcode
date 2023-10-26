package greedy.can_place_flowers;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/can-place-flowers/">605. 种花问题</a>
 * @since 2023-09-29 19:29
 */
class Solution {
    public boolean canPlaceFlowers(int[] flowerbed, int n) {
        var canPlant = 0;
        for (int i = 0, j = 0; j < flowerbed.length; j++) {
            while (j < flowerbed.length && flowerbed[j] == 0) {
                j++;
            }
            var len = j - i;
            if (i == 0) {
                len++;
            }
            if (j == flowerbed.length) {
                len++;
            }
            i = j + 1;
            canPlant += (len - 1) / 2;
        }
        return canPlant >= n;
    }
}

package string.bulls_and_cows;

import java.util.HashMap;
import java.util.Map;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/bulls-and-cows/">299. 猜数字游戏</a>
 * @since 2024-03-10 19:21
 */
class Solution {
    public String getHint(String secret, String guess) {
        int a = 0;
        int b = 0;
        int[] cnt = new int[10];
        for (int i = 0; i < secret.length(); i++) {
            char c1 = secret.charAt(i);
            char c2 = guess.charAt(i);
            if (c1 == c2) {
                a++;
            } else {
                if (cnt[c1 - '0'] < 0) {
                    b++;
                }
                if (cnt[c2 - '0'] > 0) {
                    b++;
                }
                cnt[c1 - '0']++;
                cnt[c2 - '0']--;
            }
        }
        return a + "A" + b + "B";
    }
}

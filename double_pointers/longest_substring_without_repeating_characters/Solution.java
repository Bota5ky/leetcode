package double_pointers.longest_substring_without_repeating_characters;

import java.util.HashMap;
import java.util.Map;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/longest-substring-without-repeating-characters/">3. 无重复字符的最长子串</a>
 * @link <a href="https://leetcode.cn/problems/wtcaE1/">LCR 016. 无重复字符的最长子串</a>
 * @link <a href="https://leetcode.cn/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/">LCR 167. 招式拆解 I</a>
 * @since 2023-11-03 21:00
 */
class Solution {
    public int lengthOfLongestSubstring(String s) {
        if (s.isEmpty()) {
            return 0;
        }
        Map<Character, Integer> map = new HashMap<>();
        int maxLen = 1;
        for (int i = 0, j = 0; j < s.length(); j++) {
            if (map.containsKey(s.charAt(j))) {
                i = Math.max(i, 1 + map.get(s.charAt(j)));
            }
            map.put(s.charAt(j), j);
            maxLen = Math.max(maxLen, j - i + 1);
        }
        return maxLen;
    }
}

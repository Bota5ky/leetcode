package javasolutions;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/valid-anagram/">242. 有效的字母异位词</a>
 * @since 2023-09-07 13:38
 */
class Problem242 {
    public boolean isAnagram(String s, String t) {
        int[] cnt = new int[26];
        for (int i = 0; i < s.length(); i++) {
            cnt[s.charAt(i) - 'a']++;
        }
        for (int i = 0; i < t.length(); i++) {
            cnt[t.charAt(i) - 'a']--;
        }
        for (int i = 0; i < cnt.length; i++) {
            if (cnt[i] != 0) {
                return false;
            }
        }
        return true;
    }
}

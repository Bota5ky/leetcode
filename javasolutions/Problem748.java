package javasolutions;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/shortest-completing-word/">748. 最短补全词</a>
 * @since 2023-09-07 14:18
 */
class Problem748 {
    public String shortestCompletingWord(String licensePlate, String[] words) {
        licensePlate = licensePlate.toLowerCase();
        int[] plateCnt = count(licensePlate);
        String ans = null;
        for (String word : words) {
            int[] wordCnt = count(word);
            if (compare(plateCnt, wordCnt) && (ans == null || word.length() < ans.length())) {
                ans = word;
            }
        }
        return ans;
    }

    public int[] count(String word) {
        int[] cnt = new int[26];
        for (int i = 0; i < word.length(); i++) {
            if (word.charAt(i) >= 'a' && word.charAt(i) <= 'z') {
                cnt[word.charAt(i) - 'a']++;
            }
        }
        return cnt;
    }

    public boolean compare(int[] plateCnt, int[] wordCnt) {
        for (int i = 0; i < 26; i++) {
            if (plateCnt[i] > wordCnt[i]) {
                return false;
            }
        }
        return true;
    }
}

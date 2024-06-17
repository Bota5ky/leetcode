package stack.longest_uncommon_subsequence_ii;

import java.util.Arrays;
import java.util.Stack;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/longest-uncommon-subsequence-ii/">522. 最长特殊序列 II</a>
 * @since 2024-06-17 21:35
 */
class Solution {
    public int findLUSlength(String[] strs) {
        sortByLength(strs);
        Stack<String> stack = new Stack<>();
        String pre = null;
        for (String str : strs) {
            if (stack.isEmpty()) {
                if (str.equals(pre)) {
                    continue;
                }
                stack.push(str);
            } else {
                String top = stack.peek();
                int result = compare(top, str);
                if (result == 0) {
                    stack.pop();
                } else if (result > 0) {
                    stack.pop();
                    stack.push(str);
                } else {
                    if (str.equals(pre)) {
                        continue;
                    }
                    stack.push(str);
                }
            }
            pre = str;
        }
        return stack.isEmpty() ? -1 : stack.pop().length();
    }

    private int compare(String top, String str) {
        if (top.equals(str)) {
            return 0;
        }
        int i = 0;
        for (int j = 0; j < str.length() && i < top.length(); j++) {
            if (top.charAt(i) == str.charAt(j)) {
                i++;
            }
        }
        if (i == top.length()) {
            return 1;
        }
        return -1;
    }

    private void sortByLength(String[] strs) {
        Arrays.sort(strs, (s1, s2) -> {
            if (s2.length() != s1.length()) {
                return s1.length() - s2.length();
            }
            return s1.compareTo(s2);
        });
    }
}

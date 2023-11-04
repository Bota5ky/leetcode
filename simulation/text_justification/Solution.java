package simulation.text_justification;

import java.util.ArrayList;
import java.util.List;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/text-justification/">68. 文本左右对齐</a>
 * @since 2023-09-22 15:45
 */
class Solution {
    public List<String> fullJustify(String[] words, int maxWidth) {
        int wordsLen = 0;
        int wordsCount = 0;
        ArrayList<String> result = new ArrayList<>();
        for (int i = 0; i < words.length; ) {
            if (wordsLen + wordsCount + words[i].length() <= maxWidth) {
                wordsLen += words[i].length();
                wordsCount++;
                if (i == words.length - 1) {
                    StringBuilder sb = new StringBuilder();
                    for (int j = i - wordsCount + 1; j <= i; j++) {
                        if (j == i - wordsCount + 1) {
                            sb.append(words[j]);
                        } else {
                            sb.append(" ").append(words[j]);
                        }
                    }
                    String lastLine = sb.toString();
                    int restSpaces = maxWidth - lastLine.length();
                    result.add(lastLine + repeat(" ", restSpaces));
                }
                i++;
            } else {
                int totalSpaces = maxWidth - wordsLen;
                StringBuilder sb = new StringBuilder();
                for (int j = i - wordsCount; j < i; j++) {
                    if (j == i - wordsCount) {
                        sb.append(words[j]);
                        if (wordsCount == 1) {
                            sb.append(repeat(" ", maxWidth - words[j].length()));
                        }
                    } else {
                        int appendSpaces = totalSpaces / (wordsCount - 1);
                        if (j - (i - wordsCount) <= totalSpaces % (wordsCount - 1)) {
                            appendSpaces++;
                        }
                        sb.append(repeat(" ", appendSpaces)).append(words[j]);
                    }
                }
                result.add(sb.toString());
                wordsCount = 0;
                wordsLen = 0;
            }
        }
        return result;
    }

    private String repeat(String sample, int n) {
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < n; i++) {
            sb.append(sample);
        }
        return sb.toString();
    }
}

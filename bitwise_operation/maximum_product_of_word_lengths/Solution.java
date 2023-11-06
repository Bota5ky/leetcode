package bitwise_operation.maximum_product_of_word_lengths;

import java.util.Arrays;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/maximum-product-of-word-lengths/">318. 最大单词长度乘积</a>
 * @since 2023-11-06 09:48
 */
class Solution {
    public int maxProduct(String[] words) {
        Arrays.sort(words, (o1, o2) -> o2.length() - o1.length());
        int[] bits = new int[words.length];
        for (int i = 0; i < words.length; i++) {
            for (int j = 0; j < words[i].length(); j++) {
                char c = words[i].charAt(j);
                bits[i] = bits[i] | (1 << (c - 'a'));
            }
        }
        int max = 0;
        for (int i = 0; i < words.length - 1; i++) {
            for (int j = i + 1; j < words.length && words[i].length() * words[j].length() > max; j++) {
                if ((bits[i] & bits[j]) == 0) {
                    max = words[i].length() * words[j].length();
                    break;
                }
            }
        }
        return max;
    }
}

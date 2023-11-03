package backtrack.letter_combinations_of_a_phone_number;

import java.util.ArrayList;
import java.util.List;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/letter-combinations-of-a-phone-number/">17. 电话号码的字母组合</a>
 * @since 2023-11-03 09:35
 */
class Solution {
    public List<String> letterCombinations(String digits) {
        var combinations = new ArrayList<String>();
        if (digits.isEmpty()) {
            return combinations;
        }
        backtrack(combinations, digits, 0, new StringBuilder());
        return combinations;
    }

    private void backtrack(List<String> combinations, String digits, int p, StringBuilder combination) {
        if (p == digits.length()) {
            combinations.add(combination.toString());
            return;
        }
        String[] chars = getCharacters(digits.charAt(p));
        for (String c : chars) {
            backtrack(combinations, digits, p + 1, new StringBuilder().append(combination).append(c));
        }
    }

    private String[] getCharacters(char c) {
        return switch (c) {
            case '2' -> new String[]{"a", "b", "c"};
            case '3' -> new String[]{"d", "e", "f"};
            case '4' -> new String[]{"g", "h", "i"};
            case '5' -> new String[]{"j", "k", "l"};
            case '6' -> new String[]{"m", "n", "o"};
            case '7' -> new String[]{"p", "q", "r", "s"};
            case '8' -> new String[]{"t", "u", "v"};
            case '9' -> new String[]{"w", "x", "y", "z"};
            default -> new String[0];
        };
    }
}

package string.compare_version_numbers;

import java.util.regex.Pattern;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/compare-version-numbers/">165. 比较版本号</a>
 * @since 2024-02-07 15:34
 */
class Solution {
    public int compareVersion(String version1, String version2) {
        String[] vs1 = version1.split(Pattern.quote("."));
        String[] vs2 = version2.split(Pattern.quote("."));
        for (int i = 0; i < vs1.length || i < vs2.length; i++) {
            int a = getVersion(vs1, i);
            int b = getVersion(vs2, i);
            if (a != b) {
                return a - b > 0 ? 1 : -1;
            }
        }
        return 0;
    }

    private int getVersion(String[] vs, int i) {
        if (i >= vs.length) {
            return 0;
        }
        return Integer.parseInt(vs[i]);
    }
}

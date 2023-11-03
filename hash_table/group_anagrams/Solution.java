package hash_table.group_anagrams;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/group-anagrams/">49. 字母异位词分组</a>
 * @since 2023-11-03 18:56
 */
class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        var map = new HashMap<String, ArrayList<String>>();
        for (String str : strs) {
            var key = generateKey(str);
            var list = map.getOrDefault(key, new ArrayList<>());
            list.add(str);
            map.put(key, list);
        }
        var res = new ArrayList<List<String>>();
        map.forEach((k,v) -> res.add(v));
        return res;
    }

    private String generateKey(String str) {
        var chars = new int[26];
        for (int i = 0; i < str.length(); i++) {
            chars[str.charAt(i) - 'a']++;
        }
        var sb = new StringBuilder();
        for (int i = 0; i < chars.length; i++) {
            int count = chars[i];
            if (count > 0) {
                sb.append('a' + i);
                sb.append(count);
            }
        }
        return sb.toString();
    }
}

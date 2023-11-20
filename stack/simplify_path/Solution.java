package stack.simplify_path;

import java.util.ArrayList;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/simplify-path/">71. 简化路径</a>
 * @since 2023-11-20 22:36
 */
class Solution {
    public static String simplifyPath(String path) {
        String[] dirs = path.split("/");
        ArrayList<String> list = new ArrayList<>();
        for (String dir : dirs) {
            if (dir.equals("..")) {
                if (!list.isEmpty()) {
                    list.remove(list.size() - 1);
                }
                continue;
            }
            if (!dir.equals(".") && !dir.isEmpty()) {
                list.add(dir);
            }
        }
        StringBuilder res = new StringBuilder();
        for (String dir : list) {
            res.append("/").append(dir);
        }
        if (res.length() == 0) {
            return "/";
        }
        return res.toString();
    }
}

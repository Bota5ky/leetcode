package javasolutions;

import java.util.ArrayList;
import java.util.List;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/pascals-triangle/">118. 杨辉三角</a>
 * @since 2023-09-07 11:25
 */
class Problem118 {
    public List<List<Integer>> generate(int numRows) {
        List<List<Integer>> res = new ArrayList<List<Integer>>();
        for (int i = 0; i < numRows; i++) {
            List<Integer> row = new ArrayList<>();
            for (int j = 0; j <= i; j++) {
                if (j == 0 || j == i) {
                    row.add(1);
                } else {
                    row.add(res.get(i - 1).get(j - 1) + res.get(i - 1).get(j));
                }
            }
            res.add(row);
        }
        return res;
    }
}

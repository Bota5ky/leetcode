package recursion.hanota_lcci;

import java.util.List;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/hanota-lcci/">面试题 08.06. 汉诺塔问题</a>
 * @since 2023-11-02 10:05
 */
class Solution {
    public void hanota(List<Integer> A, List<Integer> B, List<Integer> C) {
        move(A.size(), A, B, C);
    }

    private void move(int num, List<Integer> A, List<Integer> B, List<Integer> C) {
        if (num == 1) {
            C.add(A.remove(A.size() - 1));
            return;
        }
        move(num - 1, A, C, B);
        C.add(A.remove(A.size() - 1));
        move(num - 1, B, A, C);
    }
}

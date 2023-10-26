package simulation.pass_the_pillow;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/pass-the-pillow/">2582. 递枕头</a>
 * @since 2023-09-26 09:59
 */
class Solution {
    public int passThePillow(int n, int time) {
        var round = time / (n - 1);
        var remain = time % (n - 1);
        if (round % 2 == 1) {
            return n - remain;
        }
        return 1 + remain;
    }
}

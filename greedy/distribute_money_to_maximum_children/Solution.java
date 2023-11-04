package greedy.distribute_money_to_maximum_children;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/distribute-money-to-maximum-children/">2591. 将钱分给最多的儿童</a>
 * @since 2023-09-22 10:24
 */
class Solution {
    public int distMoney(int money, int children) {
        money -= children;
        if (money < 0) {
            return -1;
        }
        int max = Math.min(money / 7, children);
        int remain = money - 7 * max;
        if (remain == 3 && max == children - 1 || max == children && remain > 0) {
            return max - 1;
        }
        return max;
    }
}

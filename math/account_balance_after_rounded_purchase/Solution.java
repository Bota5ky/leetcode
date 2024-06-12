package math.account_balance_after_rounded_purchase;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/account-balance-after-rounded-purchase/">2806. 取整购买后的账户余额</a>
 * @since 2024-06-12 19:17
 */
class Solution {
    public int accountBalanceAfterPurchase(int purchaseAmount) {
        return 100 - ((purchaseAmount + 5) / 10 * 10);
    }
}

package math.majority_element;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/majority-element/">169. 多数元素</a>
 * @link <a href="https://leetcode.cn/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/">LCR 158. 库存管理 II</a>
 * @link <a href="https://leetcode.cn/problems/find-majority-element-lcci/">面试题 17.10. 主要元素 </a>
 * @since 2023-09-15 16:30
 */
class Solution {
    public int majorityElement(int[] nums) {
        var hold = Integer.MAX_VALUE;
        var count = 0;
        for (int num : nums) {
            if (count == 0 || hold == num) {
                count++;
                hold = num;
            } else {
                count--;
            }
        }
        return hold;
    }
}

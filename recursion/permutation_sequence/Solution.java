package recursion.permutation_sequence;

import java.util.ArrayList;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/permutation-sequence/">60. 排列序列</a>
 * @since 2023-11-04 14:06
 */
class Solution {
    public String getPermutation(int n, int k) {
        ArrayList<Integer> nums = new ArrayList<>();
        for (int i = 1; i <= n; i++) {
            nums.add(i);
        }
        return permutation(nums, k);
    }

    private String permutation(ArrayList<Integer> nums, int k) {
        if (k == 1) {
            StringBuilder sb = new StringBuilder();
            for (Integer num : nums) {
                sb.append(num);
            }
            return sb.toString();
        }
        int n = nums.size() - 1;
        int leftPossibilities = helper(n);
        int pos = (k - 1) / leftPossibilities;
        int cur = nums.get(pos);
        nums.remove(pos);
        return cur + permutation(nums, k - pos * leftPossibilities);
    }

    private int helper(int num) {
        int count = 1;
        for (int i = 1; i <= num; i++) {
            count *= i;
        }
        return count;
    }
}

package double_pointers.rotate_array;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/rotate-array/">189. 轮转数组</a>
 * @since 2023-09-15 17:42
 */
class Solution {
    public void rotate(int[] nums, int k) {
        k = k % nums.length;
        for (int i = 0; i < gcd(nums.length, k); i++) {
            int cur = i;
            int pre = nums[cur];
            do {
                int next = (cur + k) % nums.length;
                int temp = nums[next];
                nums[next] = pre;
                pre = temp;
                cur = next;
            } while (cur != i);
        }
    }

    private int gcd(int x, int y) {
        return y > 0 ? gcd(y, x % y) : x;
    }
}

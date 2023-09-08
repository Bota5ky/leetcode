package common.question_201_to_300;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/move-zeroes/">283. 移动零</a>
 * @since 2023-09-07 13:40
 */
class Problem283 {
    public void moveZeroes(int[] nums) {
        for (int i = 0, j = 0; j < nums.length; ) {
            if (nums[i] != 0) {
                i++;
                j++;
                continue;
            }
            do {
                j++;
            } while (j < nums.length && nums[j] == 0);
            if (j >= nums.length) {
                break;
            }
            int temp = nums[i];
            nums[i] = nums[j];
            nums[j] = temp;
            i++;
        }
    }
}

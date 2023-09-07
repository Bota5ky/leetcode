package javasolutions;

//https://leetcode-cn.com/problems/move-zeroes/
public class leetcode283 {
    public void moveZeroes(int[] nums) {
        for (int i = 0, j = 0; j < nums.length;) {
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

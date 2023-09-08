package interview

// https://leetcode-cn.com/problems/missing-number-lcci/
func missingNumber2(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	for i := 0; i < len(nums); i++ {
		sum -= nums[i]
	}
	return sum
}

// int missingNumber(vector<int>& nums)
// {
// 	int sum = 0;
// 	for (int i = 0; i < nums.size(); i++)
// 	{
// 		sum ^= i;
// 		sum ^= nums[i];
// 	}
// 	sum ^= nums.size();
// 	return sum;
// }

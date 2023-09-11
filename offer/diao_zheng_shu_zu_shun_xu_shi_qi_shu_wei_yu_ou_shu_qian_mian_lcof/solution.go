package offer

// 剑指 Offer 21. 调整数组顺序使奇数位于偶数前面 https://leetcode.cn/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
func exchange(nums []int) []int {
	i := 0
	j := len(nums) - 1
	for i < j {
		for nums[i]%2 == 1 && i < j {
			i++
		}
		for nums[j]%2 == 0 && i < j {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}

package common

import "sort"

// https://leetcode-cn.com/problems/3sum/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ret [][]int
	for i := 0; i < len(nums)-2 && nums[i] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		j := i + 1
		for k := len(nums) - 1; j < k; {
			if (j > i+1 && nums[j] == nums[j-1]) || nums[j]+nums[k] < target {
				j++
			} else if (k < len(nums)-1 && nums[k] == nums[k+1]) || nums[j]+nums[k] > target {
				k--
			} else {
				ret = append(ret, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			}
		}
	}
	return ret
}

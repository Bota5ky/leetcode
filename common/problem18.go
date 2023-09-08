package common

import "sort"

// https://leetcode-cn.com/problems/4sum/
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var ret [][]int
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		tar := target - nums[i]
		temp := trSum(nums[i+1:], tar)
		for j := 0; j < len(temp); j++ {
			ret = append(ret, []int{nums[i], temp[j][0], temp[j][1], temp[j][2]})
		}
	}
	return ret
}
func trSum(nums []int, tar int) [][]int {
	var ret [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := tar - nums[i]
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

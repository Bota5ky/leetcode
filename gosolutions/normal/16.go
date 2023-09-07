package leetcode

import "sort"

//https://leetcode-cn.com/problems/3sum-closest/
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2] - target
	ret := res + target
	if res < 0 {
		res = -res
	}
	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		for k := len(nums) - 1; j < k; {
			temp := nums[i] + nums[j] + nums[k] - target
			if temp < 0 {
				temp = -temp
			}
			if res > temp {
				res = temp
				ret = nums[i] + nums[j] + nums[k]
			}
			if nums[i]+nums[j]+nums[k] < target {
				j++
			} else if nums[i]+nums[j]+nums[k] > target {
				k--
			} else {
				return target
			}
		}
	}
	return ret
}
package gosolutions

import "sort"

// https://leetcode-cn.com/problems/sort-integers-by-the-power-value/
func getKth(lo int, hi int, k int) int {
	nums := make([]int, hi-lo+1)
	for i, j := lo, 0; i <= hi; i++ {
		nums[j] = i
		j++
	}
	sort.Slice(nums, func(i, j int) bool {
		a := weight(nums[i])
		b := weight(nums[j])
		if a == b {
			return nums[i] < nums[j]
		}
		return a < b
	})
	return nums[k-1]
}

func weight(x int) int {
	cnt := 0
	for x != 1 {
		if x%2 == 0 {
			x /= 2
		} else {
			x = 3*x + 1
		}
		cnt++
	}
	return cnt
}

package common

import "sort"

// https://leetcode.cn/problems/sort-integers-by-the-number-of-1-bits/
func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		a, b := numOf1(arr[i]), numOf1(arr[j])
		if a == b {
			return arr[i] < arr[j]
		}
		return a < b
	})
	return arr
}

func numOf1(num int) int {
	res := 0
	for num != 0 {
		res += num & 1
		num >>= 1
	}
	return res
}

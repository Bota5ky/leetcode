package sort_integers_by_the_number_of_1_bits

import "sort"

// 1356. 根据数字二进制下 1 的数目排序 https://leetcode.cn/problems/sort-integers-by-the-number-of-1-bits/
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

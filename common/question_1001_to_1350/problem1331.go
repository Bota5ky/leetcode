package question_1001_to_1350

import "sort"

// https://leetcode-cn.com/problems/rank-transform-of-an-array/
func arrayRankTransform(arr []int) []int {
	ret := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		ret[i] = arr[i]
	}
	sort.Ints(arr)
	m := make(map[int]int)
	if len(arr) > 0 {
		m[arr[0]] = 1
	}
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			m[arr[i]] = m[arr[i-1]]
		} else {
			m[arr[i]] = m[arr[i-1]] + 1
		}
	}
	for i := 0; i < len(arr); i++ {
		ret[i] = m[ret[i]]
	}
	return ret
}

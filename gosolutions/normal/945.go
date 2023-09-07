package leetcode

import "sort"

//https://leetcode-cn.com/problems/minimum-increment-to-make-array-unique/
func minIncrementForUnique(A []int) int {
	sort.Ints(A)
	cnt := 0
	for i := 1; i < len(A); i++ {
		if A[i] <= A[i-1] {
			cnt += A[i-1] + 1 - A[i]
			A[i] = A[i-1] + 1
		}

	}
	return cnt
}

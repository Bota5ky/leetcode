package 4sum_ii

import "sort"

// https://leetcode.cn/problems/4sum-ii/
func fourSumCount(A []int, B []int, C []int, D []int) int {
	sort.Ints(C)
	sort.Ints(D)
	cnt := 0
	for i := 0; i < len(A); i++ {
		target := -A[i]
		for j := 0; j < len(B); j++ {
			m, n := 0, len(D)-1
			for m < len(C) && n >= 0 {
				if B[j]+C[m]+D[n] == target {
					cntC, cntD := 0, 0
					x, y := m, n
					for ; x < len(C) && C[x] == C[m]; x++ {
						cntC++
					}
					for ; y >= 0 && D[y] == D[n]; y-- {
						cntD++
					}
					cnt += cntC * cntD
					m, n = x, y
				} else if B[j]+C[m]+D[n] > target {
					n--
				} else {
					m++
				}
			}
		}
	}
	return cnt
}

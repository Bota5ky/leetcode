package hanota_lcci

// 面试题 08.06. 汉诺塔问题 https://leetcode.cn/problems/hanota-lcci/
func hanota(A []int, B []int, C []int) []int {
	h(len(A), &A, &B, &C)
	return C
}

func h(n int, A *[]int, B *[]int, C *[]int) {
	if n == 1 {
		*C = append(*C, (*A)[len(*A)-1])
		*A = (*A)[:len(*A)-1]
		return
	}
	h(n-1, A, C, B)
	*C = append(*C, (*A)[len(*A)-1])
	*A = (*A)[:len(*A)-1]
	h(n-1, B, A, C)
}

package sorted_merge_lcci

// 面试题 10.01. 合并排序的数组 https://leetcode.cn/problems/sorted-merge-lcci/
func merge(A []int, m int, B []int, n int) {
	i := m - 1
	j := n - 1
	for k := len(A) - 1; j >= 0; k-- {
		if i < 0 || B[j] >= A[i] {
			A[k] = B[j]
			j--
		} else {
			A[k] = A[i]
			i--
		}
	}
}

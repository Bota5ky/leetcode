package subarray_sums_divisible_by_k

// https://leetcode.cn/problems/subarray-sums-divisible-by-k/
func subarraysDivByK(A []int, K int) int {
	cnt := make([]int, K)
	cnt[0] = 1
	pre, res := 0, 0
	for i := 0; i < len(A); i++ {
		pre += A[i]
		res += cnt[(pre%K+K)%K]
		cnt[(pre%K+K)%K]++
	}
	return res
}

package question_751_to_1000

// https://leetcode-cn.com/problems/subarray-sums-divisible-by-k/
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

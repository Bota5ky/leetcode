package find_n_unique_integers_sum_up_to_zero

// 1304. 和为零的 N 个不同整数 https://leetcode.cn/problems/find-n-unique-integers-sum-up-to-zero/
func sumZero(n int) []int {
	var ret []int
	for i := -n / 2; i <= n/2; i++ {
		if i == 0 && n%2 == 0 {
			continue
		}
		ret = append(ret, i)
	}
	return ret
}

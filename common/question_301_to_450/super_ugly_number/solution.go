package super_ugly_number

// 313. 超级丑数 https://leetcode.cn/problems/super-ugly-number/
func nthSuperUglyNumber(n int, primes []int) int {
	k := len(primes)
	res := make([]int, n)
	res[0] = 1
	p := make([]int, k)
	for i := 1; i < n; i++ {
		min := res[p[0]] * primes[0]
		for j := 1; j < k; j++ {
			if min > res[p[j]]*primes[j] {
				min = res[p[j]] * primes[j]
			}
		}
		res[i] = min
		for j := 0; j < k; j++ {
			if min == res[p[j]]*primes[j] {
				p[j]++
			}
		}
	}
	return res[n-1]
}

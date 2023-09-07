package gosolutions

// https://leetcode-cn.com/problems/super-ugly-number/
func nthSuperUglyNumber(n int, primes []int) int {
	k := len(primes)
	res := make([]int, n)
	res[0] = 1
	p := make([]int, k)
	for i := 1; i < n; i++ {
		//loop for n-1 times
		//k个乘数比较大小，等于最小值的指针进1，u[p[]]=乘积
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

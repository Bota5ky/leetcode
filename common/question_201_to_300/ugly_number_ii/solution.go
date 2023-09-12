package ugly_number_ii

// https://leetcode.cn/problems/ugly-number-ii/
func nthUglyNumber(n int) int {
	u := make([]int, n)
	u[0] = 1
	j, k, l := 0, 0, 0
	for i := 1; i < n; i++ {
		a := u[j] * 2
		b := u[k] * 3
		c := u[l] * 5
		if a <= b && a <= c {
			u[i] = a
			j++
		}
		if b <= a && b <= c {
			u[i] = b
			k++
		}
		if c <= a && c <= b {
			u[i] = c
			l++
		}
	}
	return u[n-1]
}

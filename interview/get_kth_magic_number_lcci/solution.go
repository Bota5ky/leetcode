package get_kth_magic_number_lcci

// https://leetcode.cn/problems/get-kth-magic-number-lcci/
func getKthMagicNumber(k int) int {
	ret := make([]int, k)
	ret[0] = 1
	x, y, z := 3, 5, 7
	i, j, l := 0, 0, 0
	for n := 1; n < k; n++ {
		a := ret[i] * x
		b := ret[j] * y
		c := ret[l] * z
		if a <= b && a <= c {
			ret[n] = a
			i++
		}
		if b <= a && b <= c {
			ret[n] = b
			j++
		}
		if c <= b && c <= a {
			ret[n] = c
			l++
		}
	}
	return ret[k-1]
}

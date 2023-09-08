package interview

// https://leetcode-cn.com/problems/convert-integer-lcci/
func convertInteger(A int, B int) int {
	cnt := 0
	for i := 0; i < 32; i++ {
		if (A^B)&(1<<i) > 0 {
			cnt++
		}
	}
	return cnt
}

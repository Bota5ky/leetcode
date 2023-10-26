package convert_integer_lcci

// 面试题 05.06. 整数转换 https://leetcode.cn/problems/convert-integer-lcci/
func convertInteger(A int, B int) int {
	cnt := 0
	for i := 0; i < 32; i++ {
		if (A^B)&(1<<i) > 0 {
			cnt++
		}
	}
	return cnt
}

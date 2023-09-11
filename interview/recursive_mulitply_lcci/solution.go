package recursive_mulitply_lcci

// https://leetcode.cn/problems/recursive-mulitply-lcci/
func multiply(A int, B int) int {
	//2  5  101
	//2  4 + 2  1
	if B == 1 {
		return A
	}
	if B&1 == 1 {
		return multiply(A, B-1) + multiply(A, 1)
	}
	return multiply(A<<1, B>>1)
}

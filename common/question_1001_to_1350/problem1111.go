package question_1001_to_1350

// https://leetcode.cn/problems/maximum-nesting-depth-of-two-valid-parentheses-strings/
func maxDepthAfterSplit(seq string) []int {
	res := make([]int, len(seq))
	isZero := 1
	for i := 0; i < len(seq); i++ {
		ch := seq[i]
		if ch == '(' {
			isZero ^= 1
			res[i] = isZero
		} else {
			res[i] = isZero
			isZero ^= 1
		}
	}
	return res
}

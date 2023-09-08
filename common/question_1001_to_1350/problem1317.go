package question_1001_to_1350

import (
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/convert-integer-to-the-sum-of-two-no-zero-integers/
func getNoZeroIntegers(n int) []int {
	var i int
	for i = 1; ; i++ {
		if iscontainZero(i) || iscontainZero(n-i) {
		} else {
			break
		}
	}
	return []int{i, n - i}
}
func iscontainZero(n int) bool {
	return strings.Contains(strconv.Itoa(n), "0")
}

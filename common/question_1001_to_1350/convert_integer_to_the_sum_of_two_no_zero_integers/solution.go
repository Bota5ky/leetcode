package convert_integer_to_the_sum_of_two_no_zero_integers

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/convert-integer-to-the-sum-of-two-no-zero-integers/
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

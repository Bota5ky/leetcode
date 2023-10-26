package convert_integer_to_the_sum_of_two_no_zero_integers

import (
	"strconv"
	"strings"
)

// 1317. 将整数转换为两个无零整数的和 https://leetcode.cn/problems/convert-integer-to-the-sum-of-two-no-zero-integers/
func getNoZeroIntegers(n int) []int {
	var i int
	for i = 1; ; i++ {
		if containZero(i) || containZero(n-i) {
		} else {
			break
		}
	}
	return []int{i, n - i}
}
func containZero(n int) bool {
	return strings.Contains(strconv.Itoa(n), "0")
}

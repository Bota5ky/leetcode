package largest_number

import (
	"sort"
	"strconv"
	"strings"
)

// 179. 最大数 https://leetcode.cn/problems/largest-number/
func largestNumber(nums []int) string {
	numbers := make([]string, len(nums))
	for i := range numbers {
		numbers[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(numbers, func(i, j int) bool {
		return strings.Compare(numbers[i]+numbers[j], numbers[j]+numbers[i]) > 0
	})
	str := strings.Join(numbers, "")
	if str[0] == '0' {
		return "0"
	}
	return str
}

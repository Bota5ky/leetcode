package largest_number

import (
	"sort"
	"strconv"
	"strings"
)

// 和offer45类似
// https://leetcode.cn/problems/largest-number/
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

/* func largestNumber(nums []int) string {
	var strNums []string
	for i := 0; i < len(nums); i++ {
		strNums = append(strNums, strconv.Itoa(nums[i]))
		for j := len(strNums) - 2; j >= 0; j-- {
			if cmpStrNums2(strNums[j], strNums[j+1]) {
				strNums[j], strNums[j+1] = strNums[j+1], strNums[j]
			}
		}
	}
	var ret string
	for i := 0; i < len(strNums); i++ {
		ret += strNums[i]
		if i == 0 && ret == "0" {
			break
		}
	}
	return ret
}

func cmpStrNums2(a string, b string) bool {
	m := a + b
	n := b + a
	for i := 0; i < len(m); i++ {
		if m[i] < n[i] {
			return true
		} else if m[i] > n[i] {
			return false
		}
	}
	return false
}
*/

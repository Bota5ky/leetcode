package leetcode

import "strconv"

//https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/
func minNumber(nums []int) string {
	var strNums []string
	for i := 0; i < len(nums); i++ {
		strNums = append(strNums, strconv.Itoa(nums[i]))
		for j := len(strNums) - 2; j >= 0; j-- {
			if cmpStrNums(strNums[j], strNums[j+1]) {
				strNums[j], strNums[j+1] = strNums[j+1], strNums[j]
			}
		}
	}
	var ret string
	for i := 0; i < len(strNums); i++ {
		ret += strNums[i]
	}
	return ret
}

func cmpStrNums(a string, b string) bool {
	m := a + b
	n := b + a
	for i := 0; i < len(m); i++ {
		if m[i] > n[i] {
			return true
		} else if m[i] < n[i] {
			return false
		}
	}
	return false
}

package common

import "strconv"

// https://leetcode.cn/problems/evaluate-reverse-polish-notation/
func evalRPN(tokens []string) int {
	var nums []int
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			switch tokens[i] {
			case "+":
				nums[len(nums)-2] += nums[len(nums)-1]
			case "-":
				nums[len(nums)-2] -= nums[len(nums)-1]
			case "*":
				nums[len(nums)-2] *= nums[len(nums)-1]
			case "/":
				nums[len(nums)-2] /= nums[len(nums)-1]
			}
			nums = nums[:len(nums)-1]
		} else {
			num, _ := strconv.Atoi(tokens[i])
			nums = append(nums, num)
		}
	}
	return nums[0]
}

package basic_calculator_ii

// https://leetcode.cn/problems/basic-calculator-ii/
func calculate(s string) int {
	var nums []int
	var temp int
	c := byte('+')
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			temp = temp*10 + int(s[i]-'0')
		}
		if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' || i == len(s)-1 {
			switch c {
			case '+':
				nums = append(nums, temp)
			case '-':
				nums = append(nums, -temp)
			case '*':
				nums[len(nums)-1] *= temp
			case '/':
				nums[len(nums)-1] /= temp
			}
			c = s[i]
			temp = 0
		}
	}
	res := 0
	for i := 0; i < len(nums); i++ {
		res += nums[i]
	}
	return res
}

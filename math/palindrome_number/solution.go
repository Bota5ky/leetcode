package palindrome_number

// 9. 回文数 https://leetcode.cn/problems/palindrome-number/
func isPalindromeNumber(x int) bool {
	if x < 0 {
		return false
	}
	var nums []int
	for x != 0 {
		nums = append(nums, x%10)
		x /= 10
	}
	for i, j := 0, len(nums)-1; i < j; {
		if nums[i] != nums[j] {
			return false
		}
		i++
		j--
	}
	return true
}

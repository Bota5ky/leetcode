package leetcode

//https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/
func missingNumber(nums []int) int {
	i := 0
	j := len(nums) - 1
	for i <= j {
		if nums[(i+j)/2] != (i+j)/2 {
			j = (i + j) / 2
		} else {
			if nums[j] == j {
				return j + 1
			}
			i = (i+j)/2 + 1
		}
		if i == j {
			break
		}
	}
	return i
}

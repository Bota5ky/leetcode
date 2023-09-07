package leetcode

//https://leetcode-cn.com/problems/trapping-rain-water/
func trap(height []int) int {
	left, right, leftmax, rightmax, res := 0, len(height)-1, 0, 0, 0
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftmax {
				leftmax = height[left]
			} else {
				res += leftmax - height[left]
			}
			left++
		} else {
			if height[right] >= rightmax {
				rightmax = height[right]
			} else {
				res += rightmax - height[right]
			}
			right--
		}
	}
	return res
}

package trapping_rain_water

// 42. 接雨水 https://leetcode.cn/problems/trapping-rain-water/
func trap(height []int) int {
	left, right, leftMax, rightMax, res := 0, len(height)-1, 0, 0, 0
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				res += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				res += rightMax - height[right]
			}
			right--
		}
	}
	return res
}

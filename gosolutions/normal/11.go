package leetcode

//https://leetcode-cn.com/problems/container-with-most-water/
func maxArea(height []int) int {
	max := 0
	for i, j := 0, len(height)-1; i < j; {
		area := (j - i) * min(height[i], height[j])
		if max < area {
			max = area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package gosolutions

// https://leetcode-cn.com/problems/binary-search/
func search(nums []int, target int) int {
	head := 0
	rear := len(nums) - 1
	for head < rear {
		mid := (head + rear) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			head = mid + 1
		} else if nums[mid] > target {
			rear = mid - 1
		}
	}
	return -1
}

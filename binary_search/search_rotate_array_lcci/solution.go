package search_rotate_array_lcci

// 面试题 10.03. 搜索旋转数组 https://leetcode.cn/problems/search-rotate-array-lcci/
func search(arr []int, target int) int {
	i, j := 0, len(arr)-1
	for i <= j {
		for j > 0 && arr[j-1] == arr[j] {
			j--
		}
		if arr[i] == target {
			return i
		}
		for i < len(arr)-1 && arr[i+1] == arr[i] {
			i++
		}
		mid := (i + j) / 2
		if arr[mid] >= arr[i] { //在左半段
			if target >= arr[i] && target <= arr[mid] {
				j = mid
			} else {
				i = mid + 1
			}
		} else {
			if target <= arr[j] && target > arr[mid] {
				i = mid + 1
			} else {
				j = mid
			}
		}
	}
	return -1
}

package kth_smallest_element_in_a_sorted_matrix

// 378. 有序矩阵中第 K 小的元素 https://leetcode.cn/problems/kth-smallest-element-in-a-sorted-matrix/
func kthSmallest(matrix [][]int, k int) int {
	nums := make([][]int, len(matrix))
	for i := 0; i < len(nums); i++ {
		nums[i] = make([]int, 2)
		nums[i] = []int{i, 0}
	}
	heapify(nums, matrix)
	for i := 0; i < k-1; i++ {
		if nums[0][1] < len(matrix[0])-1 {
			nums[0][1]++
		} else {
			nums = nums[1:]
		}
		heapify(nums, matrix)
	}
	return matrix[nums[0][0]][nums[0][1]]
}

func heapify(nums [][]int, matrix [][]int) {
	last := len(nums)/2 - 1
	for i := last; i >= 0; i-- {
		left := 2*i + 1
		right := 2*i + 2
		min := i
		if left < len(nums) && matrix[nums[min][0]][nums[min][1]] > matrix[nums[left][0]][nums[left][1]] {
			min = left
		}
		if right < len(nums) && matrix[nums[min][0]][nums[min][1]] > matrix[nums[right][0]][nums[right][1]] {
			min = right
		}
		nums[min], nums[i] = nums[i], nums[min]
	}
}

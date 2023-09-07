package gosolutions

// https://leetcode-cn.com/problems/smallest-k-lcci/
// 和offer40相同，采用部分快排
func smallestK(arr []int, k int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	quicksort(arr, k)
	return arr[:k]
}

func quicksort(nums []int, k int) {
	pivot := nums[0]
	i, j := 0, len(nums)-1
	for i < j {
		for ; i < j && nums[j] >= pivot; j-- {
		}
		nums[i] = nums[j]
		for ; i < j && nums[i] <= pivot; i++ {
		}
		nums[j] = nums[i]
	}
	nums[i] = pivot
	if i == k || i == k-1 {
		return
	} else if i > k {
		quicksort(nums[:i], k)
	} else {
		quicksort(nums[i+1:], k-i-1)
	}
}

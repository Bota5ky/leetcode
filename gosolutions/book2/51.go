package leetcode

func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

//归并排序
//https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
func mergeSort(nums []int, head int, rear int) int {
	if head >= rear {
		return 0
	}
	var count int
	if rear-head > 1 {
		count = mergeSort(nums, head, (head+rear)/2) + mergeSort(nums, (head+rear)/2+1, rear)
	}
	temp := make([]int, rear+1-head)
	copy(temp, nums[head:rear+1])
	j := len(temp) - 1 //后序列
	i := j / 2         //前序列
	mid := i
	k := rear
	for i >= 0 || j > mid {
		if j <= mid || (i >= 0 && temp[i] > temp[j]) {
			count += j - mid
			nums[k] = temp[i]
			i--
			k--
		} else {
			nums[k] = temp[j]
			j--
			k--
		}
	}
	return count
}

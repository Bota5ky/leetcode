package shu_zu_zhong_de_ni_xu_dui_lcof

// LCR 170. 交易逆序对的总数 https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

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
	j := len(temp) - 1 // 后序列
	i := j / 2         // 前序列
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

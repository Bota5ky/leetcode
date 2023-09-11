package zui_xiao_de_kge_shu_lcof

// 剑指 Offer 40. 最小的k个数 https://leetcode.cn/problems/zui-xiao-de-kge-shu-lcof/
func getLeastNumbers(arr []int, k int) []int {
	for i := k; i > 0; i-- {
		for j := len(arr) - 2; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr[:k]
}

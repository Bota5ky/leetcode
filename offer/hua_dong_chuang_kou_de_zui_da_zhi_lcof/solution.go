package hua_dong_chuang_kou_de_zui_da_zhi_lcof

// 剑指 Offer 59 - I. 滑动窗口的最大值 https://leetcode.cn/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/
// 239. 滑动窗口最大值 https://leetcode.cn/problems/sliding-window-maximum/
func maxSlidingWindow(nums []int, k int) []int {
	//双端队列
	var q, res []int
	for key, val := range nums {
		//比val小的队列q内的值删除（最大就删除所有）
		//key大于q[0]时删除q[0]
		//val比q末尾的还小就加入队列
		//key>=k时装入res[]
		if len(q) > 0 && q[0] <= key-k {
			q = q[1:]
		}
		for len(q) > 0 && nums[q[len(q)-1]] < val {
			q = q[:len(q)-1]
		}
		if len(q) == 0 || len(q) > 0 && nums[q[len(q)-1]] >= val {
			q = append(q, key)
		}
		if key >= k-1 {
			res = append(res, nums[q[0]])
		}
	}
	return res
}

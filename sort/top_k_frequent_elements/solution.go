package top_k_frequent_elements

// 347. 前 K 个高频元素 https://leetcode.cn/problems/top-k-frequent-elements/
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	cnt := make([]int, k)
	ret := make([]int, k)
	for _, v := range nums {
		m[v]++
	}
	for _, v := range m {
		if v > cnt[0] {
			cnt[0] = v
			//insert sort
			insertSort(cnt)
		}
	}
	i := 0
	for c, v := range m {
		if v >= cnt[0] {
			ret[i] = c
			i++
		}
		if i == k {
			break
		}
	}
	return ret
}

func insertSort(cnt []int) {
	temp := cnt[0]
	var i int
	for i = 1; i < len(cnt) && temp > cnt[i]; i++ {
		cnt[i-1] = cnt[i]
	}
	cnt[i-1] = temp
}

package shu_zu_zhong_shu_zi_chu_xian_de_ci_shu_lcof

// LCR 177. 撞色搭配 https://leetcode.cn/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/
func sockCollocation(nums []int) []int {
	res := 0
	//全员异或
	for _, v := range nums {
		res ^= v
	}
	k := 0 // 第K位是1
	for res&1 != 1 {
		res >>= 1
		k++
	}
	var a, b int
	for _, v := range nums {
		if v>>k&1 == 1 {
			a ^= v
		} else {
			b ^= v
		}
	}
	return []int{a, b}
}

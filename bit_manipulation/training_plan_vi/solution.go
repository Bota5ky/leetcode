package shu_zu_zhong_shu_zi_chu_xian_de_ci_shu_ii_lcof

// LCR 178. 训练计划 VI https://leetcode.cn/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/
func singleNumber(nums []int) int {
	res := make([]int, 32)
	for _, v := range nums {
		for i := 0; v > 0; i++ {
			res[i] += v & 1
			res[i] %= 3
			v >>= 1
		}
	}
	ret := 0
	for i := 31; i >= 0; i-- {
		ret <<= 1
		ret += res[i]
	}
	return ret
}

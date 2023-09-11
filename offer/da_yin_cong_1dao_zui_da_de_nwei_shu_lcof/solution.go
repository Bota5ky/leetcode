package da_yin_cong_1dao_zui_da_de_nwei_shu_lcof

// 剑指 Offer 17. 打印从1到最大的n位数 https://leetcode.cn/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/
func printNumbers(n int) []int {
	maximum := 1
	for i := 0; i < n; i++ {
		maximum *= 10
	}
	ret := make([]int, maximum-1)
	for i := 0; i < maximum-1; i++ {
		ret[i] = i + 1
	}
	return ret
}

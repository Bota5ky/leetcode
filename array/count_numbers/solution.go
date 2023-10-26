package da_yin_cong_1dao_zui_da_de_nwei_shu_lcof

// LCR 135. 报数 https://leetcode.cn/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/
func countNumbers(n int) []int {
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

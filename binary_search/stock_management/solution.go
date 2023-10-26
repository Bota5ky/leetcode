package xuan_zhuan_shu_zu_de_zui_xiao_shu_zi_lcof

// LCR 128. 库存管理 I https://leetcode.cn/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/
func stockManagement(numbers []int) int {
	i := 0
	j := len(numbers) - 1
	for i < j {
		m := (i + j) / 2
		if numbers[m] > numbers[j] {
			i = m + 1
		} else if numbers[m] < numbers[j] {
			j = m
		} else {
			j--
		}
	}
	return numbers[i]
}

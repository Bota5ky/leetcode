package he_wei_sde_lian_xu_zheng_shu_xu_lie_lcof

// LCR 180. 文件组合 https://leetcode.cn/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/
func fileCombination(target int) [][]int {
	var ret [][]int
	i, j := 1, 2
	sum := 3
	column := []int{i, j}
	for i <= target/2 {
		if sum > target {
			sum -= i
			i++
			column = column[1:]
		} else if sum < target {
			j++
			sum += j
			column = append(column, j)
		} else {
			ret = append(ret, column)
			j++
			sum += j - i
			i++
			column = column[1:]
			column = append(column, j)
		}
	}
	return ret
}

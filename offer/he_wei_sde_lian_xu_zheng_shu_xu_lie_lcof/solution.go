package he_wei_sde_lian_xu_zheng_shu_xu_lie_lcof

// 剑指 Offer 57 - II. 和为s的连续正数序列 https://leetcode.cn/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/
func findContinuousSequence(target int) [][]int {
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
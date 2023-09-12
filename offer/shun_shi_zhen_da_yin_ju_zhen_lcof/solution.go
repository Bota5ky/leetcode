package shun_shi_zhen_da_yin_ju_zhen_lcof

// 剑指 Offer 29. 顺时针打印矩阵 https://leetcode.cn/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/
func spiralOrder(matrix [][]int) []int {
	var ret []int
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return ret
	}
	a := 0
	b := len(matrix) - 1
	c := 0
	d := len(matrix[0]) - 1
	for {
		for i := c; i <= d; i++ {
			ret = append(ret, matrix[a][i])
		}
		a++
		if a > b || c > d {
			break
		}
		for i := a; i <= b; i++ {
			ret = append(ret, matrix[i][d])
		}
		d--
		if a > b || c > d {
			break
		}
		for i := d; i >= c; i-- {
			ret = append(ret, matrix[b][i])
		}
		b--
		if a > b || c > d {
			break
		}
		for i := b; i >= a; i-- {
			ret = append(ret, matrix[i][c])
		}
		c++
		if a > b || c > d {
			break
		}
	}
	return ret
}
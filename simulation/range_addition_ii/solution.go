package range_addition_ii

// 598. 范围求和 II https://leetcode.cn/problems/range-addition-ii/
func maxCount(m int, n int, ops [][]int) int {
	var minx, miny int = 40000, 40000
	for _, v := range ops {
		if v[0] < minx {
			minx = v[0]
		}
		if v[1] < miny {
			miny = v[1]
		}
	}
	return minx * miny
}

package leetcode

//https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/
func movingCount(m int, n int, k int) int {
	passed := make(map[[2]int]bool)
	return vaildcnts(0, 0, k, m, n, passed)
}

func vaildcnts(x int, y int, k int, m int, n int, passed map[[2]int]bool) int {
	if x < 0 || x >= m || y < 0 || y >= n || passed[[2]int{x, y}] == true {
		return 0
	}
	if compare(x, y, k) == false {
		return 0
	}
	passed[[2]int{x, y}] = true
	return 1 + vaildcnts(x-1, y, k, m, n, passed) + vaildcnts(x+1, y, k, m, n, passed) + vaildcnts(x, y-1, k, m, n, passed) + vaildcnts(x, y+1, k, m, n, passed)
}

func compare(x int, y int, k int) bool {
	var m, n int
	for x > 0 {
		m += x % 10
		x /= 10
	}
	for y > 0 {
		n += y % 10
		y /= 10
	}
	if m+n > k {
		return false
	}
	return true
}

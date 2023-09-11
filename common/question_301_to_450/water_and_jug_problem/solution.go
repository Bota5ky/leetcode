package question_301_to_450

// https://leetcode.cn/problems/water-and-jug-problem/
func canMeasureWater(x int, y int, z int) bool {
	if x+y < z {
		return false
	}
	if x == 0 || y == 0 {
		return z == 0 || x+y == z
	}
	return z%gcd(x, y) == 0
}

// x/y=a...b 辗转相除法
func gcd(x, y int) int {
	if x < y {
		y, x = x, y
	}
	for x%y != 0 {
		b := x % y
		x = y
		y = b
	}
	return y
}

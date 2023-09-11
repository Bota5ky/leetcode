package offer

// https://leetcode.cn/problems/qiu-12n-lcof/
func sumNums(n int) int {
	sum := 0
	num(n, &sum)
	return sum
}

func num(n int, sum *int) bool {
	*sum += n
	return n > 0 && num(n-1, sum)
}

// func sumNums(n int) int {
// 	return int(math.Pow(float64(n), 2) >> 1)
// }

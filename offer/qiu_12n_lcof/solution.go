package qiu_12n_lcof

// 剑指 Offer 64. 求1+2+…+n https://leetcode.cn/problems/qiu-12n-lcof/
func sumNums(n int) int {
	sum := 0
	num(n, &sum)
	return sum
}

func num(n int, sum *int) bool {
	*sum += n
	return n > 0 && num(n-1, sum)
}

package qiu_12n_lcof

// LCR 189. 设计机械累加器 https://leetcode.cn/problems/qiu-12n-lcof/
func mechanicalAccumulator(n int) int {
	sum := 0
	num(n, &sum)
	return sum
}

func num(n int, sum *int) bool {
	*sum += n
	return n > 0 && num(n-1, sum)
}

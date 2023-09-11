package offer

// https://leetcode.cn/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/
func findNthDigit(n int) int {
	if n < 10 {
		return n
	}
	n -= 10
	nums := 90
	i := 2
	for n >= nums*i {
		n -= nums * i
		nums *= 10
		i++
	}
	j := nums/9 + n/i
	k := i - n%i //1346  0123
	for x := 1; x < k; x++ {
		j /= 10
	}
	return j % 10
}

//1位数  0~9       10个    10
//2位数  10~99     90个    180
//3位数  100~999   900个   2700
//4位数  1000~9999 9000个

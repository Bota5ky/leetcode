package offer

// f[i]=f[i-1]+g(i,i-1)*f[i-2]
// https://leetcode.cn/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/
func translateNum(num int) int {
	f := []int{1, 1}
	last := num % 10
	num /= 10
	for i := 1; num > 0; i++ {
		fi := f[i] + g(num%10, last)*f[i-1]
		last = num % 10
		num /= 10
		f = append(f, fi)
	}
	return f[len(f)-1]
}

func g(a int, b int) int {
	if a != 0 && a*10+b < 26 {
		return 1
	}
	return 0
}

package number_of_digit_one

// 233. 数字 1 的个数 https://leetcode.cn/problems/number-of-digit-one/
func countDigitOne(n int) int {
	f := [11]int{0, 1, 20, 300, 4000, 50000, 600000, 7000000, 80000000, 900000000, 1000000000}
	j := 1
	i := 0
	num := 0
	ret := 0
	for n > 0 {
		//当前位数所包含的1  千位  1048-1000+1
		num += n % 10 * j
		present := num - j + 1
		if present > j {
			present = j
		} else if present < 0 {
			present = 0
		}
		j *= 10
		ret += present + n%10*f[i]
		n /= 10
		i++
	}
	return ret
}

package string_to_integer_atoi

// 8. 字符串转换整数 (atoi) https://leetcode.cn/problems/string-to-integer-atoi/
// LCR 192. 把字符串转换成整数 (atoi) https://leetcode.cn/problems/ba-zi-fu-chuan-zhuan-huan-cheng-zheng-shu-lcof/
func myAtoi(str string) int {
	res, i, flag := 0, 0, 1
	//去空格
	for i = 0; i < len(str) && str[i] == ' '; i++ {
	}
	if i < len(str) {
		if str[i] == '-' {
			flag = -1
			i++
		} else if str[i] == '+' {
			i++
		}
	}
	for ; i < len(str) && str[i] >= '0' && str[i] <= '9'; i++ {
		if res > (1<<31-1)/10 || res == (1<<31-1)/10 && int(str[i]-'0') > (1<<31-1)%10 {
			return 1<<31 - 1
		} else if res < -1<<31/10 || res == -1<<31/10 && -1*int(str[i]-'0') < -1<<31%10 {
			return -1 << 31
		}
		res = res*10 + flag*int(str[i]-'0')
	}
	return res
}

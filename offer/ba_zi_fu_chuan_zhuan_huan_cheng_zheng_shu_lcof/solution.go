package offer

// 剑指 Offer 67. 把字符串转换成整数 https://leetcode.cn/problems/ba-zi-fu-chuan-zhuan-huan-cheng-zheng-shu-lcof/
func strToInt(str string) int {
	//10位数
	res := 0
	flag := 1
	i := 0
	//获取符号
	for i = 0; i < len(str); i++ {
		if str[i] == ' ' {
			continue
		} else if str[i] == '+' {
			i++
			break
		} else if str[i] == '-' {
			flag = -1
			i++
			break
		} else {
			break
		}
	}
	for ; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			break
		}
		if res > (1<<31-1)/10 || res == (1<<31-1)/10 && flag*int(str[i]-'0') > (1<<31-1)%10 {
			return 1<<31 - 1
		} else if res < (-1<<31)/10 || res == (-1<<31)/10 && flag*int(str[i]-'0') < (-1<<31)%10 {
			return -1 << 31
		}
		res = res*10 + flag*int(str[i]-'0')
	}
	return res
}

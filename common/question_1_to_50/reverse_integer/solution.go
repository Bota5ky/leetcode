package reverse_integer

// 7. 整数反转 https://leetcode.cn/problems/reverse-integer/
func reverse(x int) int {
	res := 0
	for x != 0 {
		if res > (1<<31-1)/10 || res == (1<<31-1)/10 && x > (1<<31-1)%10 {
			return 0
		}
		if res < (-1<<31)/10 || res == (-1<<31)/10 && x < (-1<<31)%10 {
			return 0
		}
		res = res*10 + x%10
		x = x / 10
	}
	return res
}

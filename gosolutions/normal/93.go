package leetcode

import "strconv"

//https://leetcode-cn.com/problems/restore-ip-addresses/
func restoreIPAddresses(s string) []string {
	var pre string
	var ret []string
	addr(s, pre, 4, &ret)
	return ret
}

func addr(s, pre string, n int, ret *[]string) {
	if len(s) < n || len(s) > 3*n {
		return
	}
	if n == 1 {
		if len(s) > 1 && s[0] == '0' {	//多位数首位不能为0
			return
		}
		num, err := strconv.Atoi(s)
		if err == nil && num < 256 {
			pre += s
			*ret = append(*ret, pre)
		}
		return
	}
	for i := 1; i <= 3 && i < len(s); i++ {	//不能超出len(s)
		if i > 1 && s[0] == '0' {
			return
		}
		num, err := strconv.Atoi(s[:i])
		if err == nil && num < 256 {
			temp := pre + s[:i] + "."
			addr(s[i:], temp, n-1, ret)
		}
	}
}

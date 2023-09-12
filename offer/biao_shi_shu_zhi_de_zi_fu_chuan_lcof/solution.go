package biao_shi_shu_zhi_de_zi_fu_chuan_lcof

// 剑指 Offer 20. 表示数值的字符串 https://leetcode.cn/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/
// 65. 有效数字 https://leetcode.cn/problems/valid-number/
func isNumber(s string) bool {
	var m, n int
	for m = 0; m < len(s) && s[m] == ' '; m++ {
	}
	for n = len(s) - 1; n >= 0 && s[n] == ' '; n-- {
	}
	h := []byte(s[m:n])
	i, j := len(h), len(h) //.和e的位置
	for k := 0; k < len(h); k++ {
		if h[k] == '.' {
			i = k
		} else if h[k] == 'e' || h[k] == 'E' {
			j = k
		}
	}
	A := h[:i]
	var B, C []byte
	if i < len(h) {
		B = h[i+1 : j]
	}
	if j < len(h) {
		C = h[j+1:]
	}
	deleteSign(A)
	deleteSign(C)
	if len(A) == 0 && len(B) == 0 {
		return false
	}
	if j < len(h) && len(C) == 0 {
		return false
	}
	if judge(A) || judge(B) || judge(C) {
		return false
	}
	return true
}

func deleteSign(s []byte) {
	if len(s) != 0 && (s[0] == '+' || s[0] == '-') {
		s = s[1:]
	}
}

func judge(s []byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return true
		}
	}
	return false
}

// A[.B][eC] A\C都有可能带正负号 都是整数 AB不能同时为空 有指数符号C不能为空

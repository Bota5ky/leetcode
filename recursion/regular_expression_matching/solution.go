package regular_expression_matching

// 10. 正则表达式匹配 https://leetcode.cn/problems/regular-expression-matching/
// LCR 137. 模糊搜索验证 https://leetcode.cn/problems/zheng-ze-biao-da-shi-pi-pei-lcof/
func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	//首元素相同
	if len(p) > 1 && p[1] == '*' {
		return isMatch(s, p[2:]) || (firstMatch(s, p) && isMatch(s[1:], p)) //出现0次   和出现了的情况
	}
	//一般情况
	return firstMatch(s, p) && isMatch(s[1:], p[1:])
}

func firstMatch(s string, p string) bool {
	if len(s) != 0 && (s[0] == p[0] || p[0] == '.') {
		return true
	}
	return false
}

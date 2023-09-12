package valid_parentheses

// 20. 有效的括号 https://leetcode.cn/problems/valid-parentheses/
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '[' && s[i+1] == ']' || s[i] == '(' && s[i+1] == ')' || s[i] == '{' && s[i+1] == '}' {
			return isValid(s[:i] + s[i+2:])
		}
	}
	return false
}

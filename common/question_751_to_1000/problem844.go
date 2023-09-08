package question_751_to_1000

// https://leetcode-cn.com/problems/backspace-string-compare/
func backspaceCompare(S string, T string) bool {
	dealStr(&S)
	dealStr(&T)
	return S == T
}
func dealStr(s *string) {
	for i := 0; i < len(*s); i++ {
		if (*s)[i] == '#' {
			temp := i - 1
			if temp < 0 {
				temp = 0
			}
			*s = (*s)[:temp] + (*s)[i+1:]
			i -= 2
			if i < 0 {
				i = -1
			}
		}
	}
}

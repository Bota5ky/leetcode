package reverse_string

// 344. 反转字符串 https://leetcode.cn/problems/reverse-string/
func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

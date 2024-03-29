package greatest_common_divisor_of_strings

// 1071. 字符串的最大公因子 https://leetcode.cn/problems/greatest-common-divisor-of-strings/
func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	for {
		diff := len(str1) - len(str2)
		if diff == 0 {
			return str1
		} else if diff > 0 {
			str1 = str1[len(str2):len(str1)]
		} else {
			str2 = str2[len(str1):len(str2)]
		}
	}
}

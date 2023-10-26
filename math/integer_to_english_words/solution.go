package integer_to_english_words

import "strconv"

// 273. 整数转换英文表示 https://leetcode.cn/problems/integer-to-english-words/
// 面试题 16.08. 整数的英语表示 https://leetcode.cn/problems/english-int-lcci/
func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	//1,234,567,891 Billion Million Thousand
	n := strconv.Itoa(num)
	var ret string
	if len(n) > 9 {
		ret += convertNum(n[0]) + "Billion "
		n = n[1:]
	}
	if len(n) > 6 {
		m := n[:len(n)-6]
		n = n[len(m):]
		if len(m) == 3 && m[0] == '0' && m[1] == '0' && m[2] == '0' {
		} else {
			ret += convertThreeNums(m) + "Million "
		}
	}
	if len(n) > 3 {
		m := n[:len(n)-3]
		n = n[len(m):]
		if len(m) == 3 && m[0] == '0' && m[1] == '0' && m[2] == '0' {
		} else {
			ret += convertThreeNums(m) + "Thousand "
		}
	}
	ret += convertThreeNums(n)
	return ret[:len(ret)-1]
}
func convertThreeNums(s string) string {
	var ret string
	numLen := len(s)
	if numLen == 3 && s[0] == '0' && s[1] == '0' && s[2] == '0' {
		return ""
	}
	i := 0
	for ; numLen > 2; numLen-- {
		if s[i] != '0' {
			ret += convertNum(s[i]) + "Hundred "
		}
		i++
	}
	for ; numLen > 1; numLen-- {
		if s[i] == '1' && s[i+1] != '0' {
			switch s[i+1] {
			case '1':
				ret += "Eleven "
			case '2':
				ret += "Twelve "
			case '3':
				ret += "Thirteen "
			case '4':
				ret += "Fourteen "
			case '5':
				ret += "Fifteen "
			case '6':
				ret += "Sixteen "
			case '7':
				ret += "Seventeen "
			case '8':
				ret += "Eighteen "
			case '9':
				ret += "Nineteen "
			}
			numLen = 0
			break
		}
		ret += tens(s[i])
		i++
	}
	for ; numLen > 0; numLen-- {
		ret += convertNum(s[i])
		i++
	}
	return ret
}
func tens(s byte) string {
	switch s {
	case '1':
		return "Ten "
	case '2':
		return "Twenty "
	case '3':
		return "Thirty "
	case '4':
		return "Forty "
	case '5':
		return "Fifty "
	case '6':
		return "Sixty "
	case '7':
		return "Seventy "
	case '8':
		return "Eighty "
	case '9':
		return "Ninety "
	}
	return ""
}
func convertNum(s byte) string {
	switch s {
	case '1':
		return "One "
	case '2':
		return "Two "
	case '3':
		return "Three "
	case '4':
		return "Four "
	case '5':
		return "Five "
	case '6':
		return "Six "
	case '7':
		return "Seven "
	case '8':
		return "Eight "
	case '9':
		return "Nine "
	}
	return ""
}

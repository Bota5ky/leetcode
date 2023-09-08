package interview

// 和273相同
// https://leetcode-cn.com/problems/english-int-lcci/
func numberToWords2(num int) string {
	//123,456,789   billion million thousand hundred 1~9 or 10~19 or **ty + 个位数
	if num == 0 {
		return "Zero"
	}
	var ret string
	cnt := 0
	for num > 0 {
		temp := num % 1000
		tempstr := dealWithThreeNums(temp)
		switch cnt {
		case 1:
			if temp > 0 {
				tempstr += " Thousand"
			} //100,000
		case 2:
			if temp > 0 {
				tempstr += " Million"
			} //1,000,000
		case 3:
			tempstr += " Billion" //1,000,000,000
		}
		ret = tempstr + ret
		cnt++
		num /= 1000
	}
	return ret[1:]
}
func dealWithThreeNums(num int) string {
	var ret string
	//百位
	if num/100 > 0 {
		ret += oneTo19(num/100) + " Hundred"
	}
	tens := num % 100
	if tens < 20 {
		ret += oneTo19(tens)
	} else {
		ret += over20(tens/10) + oneTo19(tens%10)
	}
	return ret
}
func oneTo19(num int) string {
	switch num {
	case 1:
		return " One"
	case 2:
		return " Two"
	case 3:
		return " Three"
	case 4:
		return " Four"
	case 5:
		return " Five"
	case 6:
		return " Six"
	case 7:
		return " Seven"
	case 8:
		return " Eight"
	case 9:
		return " Nine"
	case 10:
		return " Ten"
	case 11:
		return " Eleven"
	case 12:
		return " Twelve"
	case 13:
		return " Thirteen"
	case 14:
		return " Fourteen"
	case 15:
		return " Fifteen"
	case 16:
		return " Sixteen"
	case 17:
		return " Seventeen"
	case 18:
		return " Eighteen"
	case 19:
		return " Nineteen"
	}
	return ""
}

func over20(num int) string {
	switch num {
	case 2:
		return " Twenty"
	case 3:
		return " Thirty"
	case 4:
		return " Forty"
	case 5:
		return " Fifty"
	case 6:
		return " Sixty"
	case 7:
		return " Seventy"
	case 8:
		return " Eighty"
	case 9:
		return " Ninety"
	}
	return ""
}

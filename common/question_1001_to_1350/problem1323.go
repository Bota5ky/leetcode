package question_1001_to_1350

//https://leetcode.cn/problems/maximum-69-number/
import (
	"strconv"
)

func maximum69Number(num int) int {
	n := strconv.Itoa(num)
	ret := []byte(n)
	for i := 0; i < len(ret); i++ {
		if ret[i] == '6' {
			ret[i] = '9'
			break
		}
	}
	x, _ := strconv.Atoi(string(ret))
	return x
}

package leetcode

import "strconv"

//https://leetcode-cn.com/problems/fizz-buzz/
func fizzBuzz(n int) []string {
	res := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			res[i-1] = "FizzBuzz"
			continue
		}
		if i%3 == 0 {
			res[i-1] = "Fizz"
			continue
		}
		if i%5 == 0 {
			res[i-1] = "Buzz"
			continue
		}
		res[i-1] = strconv.Itoa(i)
	}
	return res
}

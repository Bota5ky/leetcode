package common

import "math"

// https://leetcode-cn.com/problems/powerful-integers/
func powerfulIntegers(x int, y int, bound int) []int {
	var ret []int
	m := make(map[int]bool)
	for i := 0; int(math.Pow(float64(x), float64(i)))+1 <= bound; i++ {
		for j := 0; ; j++ {
			temp := int(math.Pow(float64(x), float64(i))) + int(math.Pow(float64(y), float64(j)))
			if temp <= bound && m[temp] == false {
				m[temp] = true
				ret = append(ret, temp)
			}
			if temp > bound || y == 1 {
				break
			}
		}
		if x == 1 {
			break
		}
	}
	return ret
}

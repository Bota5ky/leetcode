package find_lucky_integer_in_an_array

// https://leetcode.cn/problems/find-lucky-integer-in-an-array/
func findLucky(arr []int) int {
	cnt := make(map[int]int)
	for _, v := range arr {
		cnt[v]++
	}
	res := -1
	for c, v := range cnt {
		if c == v {
			if res < c {
				res = c
			}
		}
	}
	return res
}

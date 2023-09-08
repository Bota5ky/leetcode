package common

// https://leetcode-cn.com/problems/distribute-candies-to-people/
func distributeCandies(candies int, numPeople int) []int {
	ret := make([]int, numPeople)
	j := 1
	for i := 0; candies > 0; i++ {
		if i == numPeople {
			i = 0
		}
		if candies > j {
			ret[i] += j
			candies -= j
		} else {
			ret[i] += candies
			break
		}
		j++
	}
	return ret
}

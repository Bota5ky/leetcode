package gosolutions

// https://leetcode-cn.com/problems/kids-with-the-greatest-number-of-candies/
func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	for _, v := range candies {
		if v > max {
			max = v
		}
	}
	res := make([]bool, len(candies))
	for c, v := range candies {
		if v+extraCandies >= max {
			res[c] = true
		}
	}
	return res
}

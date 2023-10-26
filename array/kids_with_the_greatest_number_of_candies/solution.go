package kids_with_the_greatest_number_of_candies

// 1431. 拥有最多糖果的孩子 https://leetcode.cn/problems/kids-with-the-greatest-number-of-candies/
func kidsWithCandies(candies []int, extraCandies int) []bool {
	maximum := 0
	for _, v := range candies {
		if v > maximum {
			maximum = v
		}
	}
	res := make([]bool, len(candies))
	for c, v := range candies {
		if v+extraCandies >= maximum {
			res[c] = true
		}
	}
	return res
}

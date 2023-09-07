package leetcode

//https://leetcode-cn.com/problems/longest-palindrome/
func longestPalindrome(s string) int {
	var sum int
	m := make(map[byte]int)
	odd := 0
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		if m[s[i]]%2 == 0 {
			sum += 2
			odd--
		} else {
			odd++
		}
	}
	if odd > 0 {
		odd = 1
	}
	return sum + odd
}

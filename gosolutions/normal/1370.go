package leetcode

//https://leetcode-cn.com/problems/increasing-decreasing-string/
func sortString(s string) string {
	cnt := make([]int, 26)
	sum := 0
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
		sum++
	}
	res := ""
	for sum > 0 {
		for i := 0; i < 26 && sum > 0; i++ {
			if cnt[i] > 0 {
				res += string(byte(i) + 'a')
				cnt[i]--
				sum--
			}
		}
		for i := 25; i >= 0 && sum > 0; i-- {
			if cnt[i] > 0 {
				res += string(byte(i) + 'a')
				cnt[i]--
				sum--
			}
		}
	}
	return res
}

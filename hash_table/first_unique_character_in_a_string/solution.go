package first_unique_character_in_a_string

// 387. 字符串中的第一个唯一字符 https://leetcode.cn/problems/first-unique-character-in-a-string/
func firstUniqChar(s string) int {
	cnt := make([]int, 26)
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if cnt[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}

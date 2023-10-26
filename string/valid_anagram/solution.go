package valid_anagram

// 242. 有效的字母异位词 https://leetcode.cn/problems/valid-anagram/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var (
		cnt = make([]int, 26)
	)
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
		cnt[t[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if cnt[i] != 0 {
			return false
		}
	}
	return true
}

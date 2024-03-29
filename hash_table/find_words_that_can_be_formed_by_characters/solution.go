package find_words_that_can_be_formed_by_characters

// 1160. 拼写单词 https://leetcode.cn/problems/find-words-that-can-be-formed-by-characters/
func countCharacters(words []string, chars string) int {
	charsCnt := make([]int, 26)
	for i := 0; i < len(chars); i++ {
		charsCnt[chars[i]-'a']++
	}
	length := 0
	for i := 0; i < len(words); i++ {
		temp := make([]int, 26)
		copy(temp, charsCnt)
		flag := true
		for j := 0; j < len(words[i]); j++ {
			temp[words[i][j]-'a']--
			if temp[words[i][j]-'a'] < 0 {
				flag = false
				break
			}
		}
		if flag {
			length += len(words[i])
		}
	}
	return length
}

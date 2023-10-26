package verifying_an_alien_dictionary

// 953. 验证外星语词典 https://leetcode.cn/problems/verifying-an-alien-dictionary/
func isAlienSorted(words []string, order string) bool {
	//录入字典序
	m := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		m[order[i]] = i
	}
	//对比
	for i := 0; i < len(words)-1; i++ {
		for j := 0; j < len(words[i]); j++ {
			if j >= len(words[i+1]) || m[words[i][j]] > m[words[i+1][j]] {
				return false
			} else if m[words[i][j]] < m[words[i+1][j]] {
				break
			}
		}
	}
	return true
}

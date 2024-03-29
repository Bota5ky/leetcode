package shortest_completing_word

import "strings"

// 748. 最短补全词 https://leetcode.cn/problems/shortest-completing-word/
func shortestCompletingWord(licensePlate string, words []string) string {
	licensePlate = strings.ToLower(licensePlate)
	m := make(map[byte]int)
	for i := 0; i < len(licensePlate); i++ {
		if licensePlate[i] >= 'a' && licensePlate[i] <= 'z' {
			m[licensePlate[i]]++
		}
	}
	minLen := 16
	var ret string
	for i := 0; i < len(words); i++ {
		if len(words[i]) >= minLen {
			continue
		}
		mark := 0
		for c, val := range m {
			if strings.Count(words[i], string(c)) < val {
				mark = 1
				break
			}
		}
		if mark == 0 {
			ret = words[i]
			minLen = len(words[i])
		}
	}
	return ret
}

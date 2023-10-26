package find_the_longest_substring_containing_vowels_in_even_counts

// 1371. 每个元音包含偶数次的最长子字符串 https://leetcode.cn/problems/find-the-longest-substring-containing-vowels-in-even-counts/
func findTheLongestSubstring(s string) int {
	cnt, max := 0, 0
	m := make(map[int]int)
	m[0] = -1
	for key := range s {
		switch s[key] {
		case 'a':
			cnt ^= 16
		case 'e':
			cnt ^= 8
		case 'i':
			cnt ^= 4
		case 'o':
			cnt ^= 2
		case 'u':
			cnt ^= 1
		}
		if val, ok := m[cnt]; ok {
			if max < key-val {
				max = key - val
			}
		} else {
			m[cnt] = key
		}
	}
	return max
}

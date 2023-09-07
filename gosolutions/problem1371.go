package gosolutions

// https://leetcode-cn.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/
func findTheLongestSubstring(s string) int {
	//aeiou 11111
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

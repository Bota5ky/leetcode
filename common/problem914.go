package common

// https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards/
func hasGroupsSizeX(deck []int) bool {
	m, maxcnt := make(map[int]int), 2
	for _, v := range deck {
		m[v]++
		if m[v] > maxcnt {
			maxcnt = m[v]
		}
	}
	for i := 2; i <= maxcnt; i++ {
		flag := true
		for _, v := range m {
			if v%i != 0 {
				flag = false
				break
			}
		}
		if flag {
			return true
		}
	}
	return false
}

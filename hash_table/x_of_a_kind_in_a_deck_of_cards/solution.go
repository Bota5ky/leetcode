package x_of_a_kind_in_a_deck_of_cards

// 914. 卡牌分组 https://leetcode.cn/problems/x-of-a-kind-in-a-deck-of-cards/
func hasGroupsSizeX(deck []int) bool {
	m, maxCnt := make(map[int]int), 2
	for _, v := range deck {
		m[v]++
		if m[v] > maxCnt {
			maxCnt = m[v]
		}
	}
	for i := 2; i <= maxCnt; i++ {
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

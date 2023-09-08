package interview

// https://leetcode-cn.com/problems/find-closest-lcci/
func findClosest(words []string, word1 string, word2 string) int {
	//string->[]int坐标位置,再用binary search
	m := make(map[string][]int)
	for key, word := range words {
		m[word] = append(m[word], key)
	}
	dis1 := m[word1]
	dis2 := m[word2]
	min := 2<<31 - 1
	for i := 0; i < len(dis1); i++ {
		for j := 0; j < len(dis2); j++ {
			if abs(dis1[i]-dis2[j]) < min {
				min = abs(dis1[i] - dis2[j])
			}
		}
	}
	return min
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

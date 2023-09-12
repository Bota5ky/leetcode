package master_mind_lcci

// 面试题 16.15. 珠玑妙算 https://leetcode.cn/problems/master-mind-lcci/
func masterMind(solution string, guess string) []int {
	x := make([]int, 4)
	y := make([]int, 4)
	same := 0
	for i := 0; i < 4; i++ {
		figureColor(x, guess[i])
		figureColor(y, solution[i])
		if guess[i] == solution[i] {
			same++
		}
	}
	diff := 0
	for i := 0; i < 4; i++ {
		if x[i] <= y[i] {
			diff += x[i]
		} else {
			diff += y[i]
		}
	}
	return []int{same, diff - same}
}

func figureColor(cnt []int, color byte) {
	switch color {
	case 'R':
		cnt[0]++
	case 'Y':
		cnt[1]++
	case 'G':
		cnt[2]++
	case 'B':
		cnt[3]++
	}
}

package leetcode

//https://leetcode-cn.com/problems/master-mind-lcci/
func masterMind(solution string, guess string) []int {
	//RYGB
	x := make([]int, 4) //guess
	y := make([]int, 4) //solution
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

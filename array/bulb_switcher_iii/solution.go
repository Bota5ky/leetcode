package bulb_switcher_iii

// 1375. 灯泡开关 III https://leetcode.cn/problems/bulb-switcher-iii/
func numTimesAllBlue(light []int) int {
	cnt := 0
	sum := 0
	for i := 0; i < len(light); i++ {
		sum += light[i] - 1
		sum -= i
		if sum == 0 {
			cnt++
		}
	}
	return cnt
}

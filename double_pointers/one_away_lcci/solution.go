package one_away_lcci

// 面试题 01.05. 一次编辑 https://leetcode.cn/problems/one-away-lcci/
func oneEditAway(first string, second string) bool {
	if len(first) == 0 {
		return len(second) <= 1
	}
	if len(second) == 0 {
		return len(first) <= 1
	}
	if first[0] == second[0] {
		return oneEditAway(first[1:], second[1:])
	} else if first[len(first)-1] == second[len(second)-1] {
		return oneEditAway(first[:len(first)-1], second[:len(second)-1])
	}
	return len(first) == 1 && len(second) == 1
}

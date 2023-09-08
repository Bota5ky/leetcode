package common

import "math"

// https://leetcode-cn.com/problems/minimum-window-substring/
func minWindow(s string, t string) string {
	ori := [256]int{}
	cnt := [256]int{}
	maxLen := math.MaxInt32
	for _, val := range t {
		ori[int(val)]++
	}
	check := func() bool {
		for key, val := range cnt {
			if val < ori[key] { //不包含
				return false
			}
		}
		return true
	}
	ansL, ansR := -1, -1
	for i, j := 0, 0; j < len(s); j++ {
		if j < len(s) && ori[int(s[j])] > 0 {
			cnt[int(s[j])]++
		}
		for check() && i <= j {
			if j-i+1 < maxLen {
				ansL, ansR = i, j+1
				maxLen = j - i + 1
			}
			if ori[int(s[i])] > 0 {
				cnt[int(s[i])]--
			}
			i++
		}
	}
	if ansR == -1 {
		return ""
	}
	return s[ansL:ansR]
}

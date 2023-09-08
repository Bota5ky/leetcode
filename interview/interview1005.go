package interview

// https://leetcode-cn.com/problems/sparse-array-search-lcci/
func findString(words []string, s string) int {
	i, j := 0, len(words)-1
	for i < j {
		mid := (i + j) / 2
		for ; mid > i; mid-- {
			if words[mid] != "" {
				break
			}
			if mid == i+1 {
				break
			}
		}
		if words[mid] < s {
			i = mid + 1
		} else if words[mid] > s {
			j = mid
		} else {
			return mid
		}
	}
	if words[i] == s {
		return i
	}
	return -1
}

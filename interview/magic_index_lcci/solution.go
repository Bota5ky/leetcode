package magic_index_lcci

// https://leetcode.cn/problems/magic-index-lcci/
func findMagicIndex(nums []int) int {
	for c, v := range nums {
		if c == v {
			return c
		}
	}
	return -1
}

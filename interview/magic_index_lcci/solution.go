package magic_index_lcci

// 面试题 08.03. 魔术索引 https://leetcode.cn/problems/magic-index-lcci/
func findMagicIndex(nums []int) int {
	for c, v := range nums {
		if c == v {
			return c
		}
	}
	return -1
}

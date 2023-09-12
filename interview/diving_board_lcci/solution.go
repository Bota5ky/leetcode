package diving_board_lcci

// 面试题 16.11. 跳水板 https://leetcode.cn/problems/diving-board-lcci/
func divingBoard(shorter int, longer int, k int) []int {
	var res []int
	if k == 0 {
		return res
	}
	if shorter == longer {
		return []int{k * shorter}
	}
	for i := k; i >= 0; i-- {
		res = append(res, shorter*i+longer*(k-i))
	}
	return res
}
